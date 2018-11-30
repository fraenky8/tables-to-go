package tablestogo

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"

	"github.com/fraenky8/tables-to-go/src/database"
	"github.com/fraenky8/tables-to-go/src/settings"
	"github.com/fraenky8/tables-to-go/src/tagger"
)

var (
	// map of Tagger used
	// key is a ascending sequence of i*2 to determine which tags to generate later
	taggers = map[int]tagger.Tagger{
		1: new(tagger.DbTag),
		2: new(tagger.StblTag),
		4: new(tagger.SQLTag),
	}

	// means that the `db`-Tag is enabled by default
	effectiveTags = 1
)

// Run is the main function to run the conversions
func Run(settings *settings.Settings) (err error) {

	createEffectiveTags(settings)

	db := database.NewDatabase(settings)

	if err = db.Connect(); err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}
	defer db.Close()

	return run(settings, db)
}

func createEffectiveTags(settings *settings.Settings) {
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

func run(settings *settings.Settings, db database.Database) (err error) {

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

func createTableStructString(settings *settings.Settings, db database.Database, table *database.Table) (string, string) {

	var structFields strings.Builder

	var isNullable bool
	var isTime bool

	for _, column := range table.Columns {

		// TODO add verbosity levels
		// if settings.Verbose {
		// 	fmt.Printf("\t> %v\r\n", column.Name)
		// }

		column.Name = strings.Title(column.Name)
		if settings.OutputFormat == "c" {
			column.Name = camelCaseString(column.Name)
		}
		columnType, isTimeType := mapDbColumnTypeToGoType(db, column)

		// ISSUE-4: if columns are part of multiple constraints
		// then the sql returns multiple rows per column name.
		// Therefore we check if we already added a column with
		// that name to the struct, if so, skip.
		if strings.Contains(structFields.String(), column.Name+" ") {
			continue
		}

		structFields.WriteString(column.Name)
		structFields.WriteString(" ")
		structFields.WriteString(columnType)
		structFields.WriteString(generateTags(db, column))
		structFields.WriteString("\n")

		// save some info for later use
		if column.IsNullable == "YES" {
			isNullable = true
		}
		if isTimeType {
			isTime = true
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

	// do imports
	if isNullable || isTime || settings.IsMastermindStructableRecorder {
		fileContent.WriteString("import (\n")

		if isNullable {
			fileContent.WriteString("\t\"database/sql\"\n")
		}

		if isTime {
			if isNullable {
				fileContent.WriteString("\t\n\"github.com/lib/pq\"\n")
			} else {
				fileContent.WriteString("\t\"time\"\n")
			}
		}

		if settings.IsMastermindStructableRecorder {
			fileContent.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
		}

		fileContent.WriteString(")\n\n")
	}

	tableName := strings.Title(settings.Prefix + table.Name + settings.Suffix)
	if settings.OutputFormat == "c" {
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

func createStructFile(path, name, content string) error {

	fileName := path + name + ".go"

	// format it
	formatedContent, err := format.Source([]byte(content))
	if err != nil {
		return fmt.Errorf("could not format file %s: %v", fileName, err)
	}

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

func mapDbColumnTypeToGoType(db database.Database, column database.Column) (goType string, isTime bool) {

	isTime = false

	if db.IsString(column) || db.IsText(column) {
		goType = "string"
		if db.IsNullable(column) {
			goType = "sql.NullString"
		}
	} else if db.IsInteger(column) {
		goType = "int"
		if db.IsNullable(column) {
			goType = "sql.NullInt64"
		}
	} else if db.IsFloat(column) {
		goType = "float64"
		if db.IsNullable(column) {
			goType = "sql.NullFloat64"
		}
	} else if db.IsTemporal(column) {
		goType = "time.Time"
		if db.IsNullable(column) {
			goType = "pq.NullTime"
		}
		isTime = true
	} else {

		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if db.IsNullable(column) {
				goType = "sql.NullBool"
			}
		default:
			goType = "sql.NullString"
		}
	}

	return goType, isTime
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
