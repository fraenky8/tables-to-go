package cli

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"

	"github.com/fraenky8/tables-to-go/pkg/config"
	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/database/mysql"
	"github.com/fraenky8/tables-to-go/pkg/database/postgresql"
	"github.com/fraenky8/tables-to-go/pkg/tagger"
)

var (
	// map of Tagger used
	// key is a ascending sequence of i*2 to determine which tags to generate later
	taggers = map[int]tagger.Tagger{
		1: new(tagger.Db),
		2: new(tagger.Mastermind),
		4: new(tagger.SQL),
	}

	// means that the `db`-Tag is enabled by default
	effectiveTags = 1

	// some strings for idiomatic go in column names
	// see https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	initialisms = []string{"ID", "JSON", "XML", "HTTP", "URL"}
)

// Run runs the transformations by creating the concrete Database by the provided settings
func Run(settings *config.Settings) (err error) {

	db, err := newDatabase(settings)
	if err != nil {
		return err
	}

	createEffectiveTags(settings)

	fmt.Printf("running for %q...\r\n", settings.DbType)

	tables, err := db.GetTables()
	if err != nil {
		return fmt.Errorf("could not get tables: %v", err)
	}

	if settings.Verbose {
		fmt.Printf("> number of tables: %v\r\n", len(tables))
	}

	if err = db.PrepareGetColumnsOfTableStmt(); err != nil {
		return fmt.Errorf("could not prepare the get-column-statement: %v", err)
	}

	for _, table := range tables {

		if settings.Verbose {
			fmt.Printf("> processing table %q\r\n", table.Name)
		}

		if err = db.GetColumnsOfTable(table); err != nil {
			return fmt.Errorf("could not get columns of table %s: %v", table.Name, err)
		}

		if settings.Verbose {
			fmt.Printf("\t> number of columns: %v\r\n", len(table.Columns))
		}

		tableName, content := createTableStructString(settings, db, table)

		if err = createStructFile(settings.OutputFilePath, tableName, content); err != nil {
			return fmt.Errorf("could not create struct file for table %s: %v", table.Name, err)
		}
	}

	fmt.Println("done!")

	return err
}

func newDatabase(settings *config.Settings) (database.Database, error) {

	gdb := database.New(settings)

	var db database.Database

	switch settings.DbType {
	case "mysql":
		db = mysql.New(gdb)
	case "pg":
		fallthrough
	default:
		db = postgresql.New(gdb)
	}

	if err := db.Connect(); err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	return db, nil
}

func createEffectiveTags(settings *config.Settings) {
	if settings.TagsNoDb {
		effectiveTags = 0
	}
	if settings.TagsMastermindStructable {
		effectiveTags |= 2
	}
	if settings.TagsMastermindStructableOnly {
		effectiveTags = 0
		effectiveTags |= 2
	}
	if settings.TagsSQL {
		effectiveTags |= 4
	}
	if settings.TagsSQLOnly {
		effectiveTags = 0
		effectiveTags |= 4
	}
	// last tag-"ONLY" wins if multiple specified
}

type columnInfo struct {
	isTime             bool
	isNullablePrimitve bool
	isNullableTime     bool
}

func (c columnInfo) hasTrue() bool {
	return c.isTime || c.isNullableTime || c.isNullablePrimitve
}

func createTableStructString(settings *config.Settings, db database.Database, table *database.Table) (string, string) {

	var structFields strings.Builder

	columnInfo := columnInfo{}
	columns := map[string]struct{}{}

	for _, column := range table.Columns {

		columnName := strings.Title(column.Name)
		if settings.OutputFormat == config.OutputFormatCamelCase {
			columnName = camelCaseString(column.Name)
			columnName = toInitialisms(columnName)
		}
		columnType, isTimeType := mapDbColumnTypeToGoType(settings, db, column)

		// ISSUE-4: if columns are part of multiple constraints
		// then the sql returns multiple rows per column name.
		// Therefore we check if we already added a column with
		// that name to the struct, if so, skip.
		if _, ok := columns[columnName]; ok {
			continue
		}
		columns[columnName] = struct{}{}

		if settings.VVerbose {
			fmt.Printf("\t\t> %v\r\n", column.Name)
		}

		structFields.WriteString(columnName)
		structFields.WriteString(" ")
		structFields.WriteString(columnType)
		structFields.WriteString(generateTags(db, column))
		structFields.WriteString("\n")

		// save some info for later use
		columnInfo.isNullablePrimitve = db.IsNullable(column) && !db.IsTemporal(column)

		// save that we saw a time type column at least once
		if isTimeType {
			columnInfo.isTime = true
			columnInfo.isNullableTime = db.IsNullable(column)
		}
	}

	if settings.IsMastermindStructableRecorder {
		structFields.WriteString("\t\nstructable.Recorder\n")
	}

	var fileContent strings.Builder

	// write header infos
	fileContent.WriteString("package ")
	fileContent.WriteString(settings.PackageName)
	fileContent.WriteString("\n\n")

	// write imports
	generateImports(&fileContent, settings, db, columnInfo)

	tableName := strings.Title(settings.Prefix + table.Name + settings.Suffix)
	if settings.OutputFormat == config.OutputFormatCamelCase {
		tableName = camelCaseString(tableName)
	}

	// write struct with fields
	fileContent.WriteString("type ")
	fileContent.WriteString(tableName)
	fileContent.WriteString(" struct {\n")
	fileContent.WriteString(structFields.String())
	fileContent.WriteString("}")

	return tableName, fileContent.String()
}

func generateImports(content *strings.Builder, settings *config.Settings, db database.Database, columnInfo columnInfo) {

	if !columnInfo.hasTrue() && !settings.IsMastermindStructableRecorder {
		return
	}

	content.WriteString("import (\n")

	if columnInfo.isNullablePrimitve && settings.IsNullTypeSQL() {
		content.WriteString("\t\"database/sql\"\n")
	}

	if settings.IsMastermindStructableRecorder {
		content.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
	}

	if columnInfo.isTime {
		if columnInfo.isNullableTime && settings.IsNullTypeSQL() {
			content.WriteString("\t\n")
			content.WriteString(db.GetDriverImportLibrary())
			content.WriteString("\n")
		} else {
			content.WriteString("\t\"time\"\n")
		}
	}

	content.WriteString(")\n\n")
}

func createStructFile(path, name, content string) error {

	fileName := path + name + ".go"

	// format it
	formatedContent, err := format.Source([]byte(content))
	if err != nil {
		return fmt.Errorf("could not format file %s: %v", fileName, err)
	}

	// fight the symptom instead of the cause - if we didnt imported anything, remove it
	formatedContent = bytes.ReplaceAll(formatedContent, []byte("\nimport ()\n"), []byte(""))

	return ioutil.WriteFile(fileName, formatedContent, 0666)
}

func generateTags(db database.Database, column database.Column) (tags string) {
	for t := 1; t <= effectiveTags; t *= 2 {
		shouldTag := effectiveTags&t > 0
		if shouldTag {
			tags += taggers[t].GenerateTag(db, column) + " "
		}
	}
	if len(tags) > 0 {
		tags = " `" + strings.TrimSpace(tags) + "`"
	}
	return tags
}

func mapDbColumnTypeToGoType(settings *config.Settings, db database.Database, column database.Column) (goType string, isTime bool) {

	isTime = false

	if db.IsString(column) || db.IsText(column) {
		goType = "string"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*string", "sql.NullString")
		}
	} else if db.IsInteger(column) {
		goType = "int"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*int", "sql.NullInt64")
		}
	} else if db.IsFloat(column) {
		goType = "float64"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*float64", "sql.NullFloat64")
		}
	} else if db.IsTemporal(column) {
		goType = "time.Time"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*time.Time", db.GetTemporalDriverDataType())
		}
		isTime = true
	} else {
		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if db.IsNullable(column) {
				goType = getNullType(settings, "*bool", "sql.NullBool")
			}
		default:
			goType = getNullType(settings, "*string", "sql.NullString")
		}
	}

	return goType, isTime
}

func getNullType(settings *config.Settings, primitive string, sql string) string {
	if settings.IsNullTypeSQL() {
		return sql
	}
	return primitive
}

func camelCaseString(s string) (cc string) {
	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		return strings.Title(s)
	}

	for _, part := range splitted {
		cc += strings.Title(strings.ToLower(part))
	}
	return cc
}

func toInitialisms(s string) string {
	for _, substr := range initialisms {
		idx := indexCaseInsensitive(s, substr)
		if idx == -1 {
			continue
		}
		toReplace := s[idx : idx+len(substr)]
		s = strings.ReplaceAll(s, toReplace, substr)
	}
	return s
}

func indexCaseInsensitive(s, substr string) int {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Index(s, substr)
}
