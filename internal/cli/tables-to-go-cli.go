package cli

import (
	"fmt"
	"strings"

	"github.com/fraenky8/tables-to-go/pkg/config"
	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
	"github.com/fraenky8/tables-to-go/pkg/tagger"
)

var (
	taggers tagger.Tagger

	// some strings for idiomatic go in column names
	// see https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	initialisms = []string{"ID", "JSON", "XML", "HTTP", "URL"}
)

// Run runs the transformations by creating the concrete Database by the provided settings
func Run(settings *config.Settings, db database.Database, out output.Writer) (err error) {

	taggers = tagger.NewTaggers(settings)

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

		err = out.Write(tableName, content)
		if err != nil {
			return fmt.Errorf("could not write struct for table %s: %v", table.Name, err)
		}
	}

	fmt.Println("done!")

	return nil
}

type columnInfo struct {
	isNullable          bool
	isTemporal          bool
	isNullablePrimitive bool
	isNullableTemporal  bool
}

func (c columnInfo) hasTrue() bool {
	return c.isNullable || c.isTemporal || c.isNullableTemporal || c.isNullablePrimitive
}

func createTableStructString(settings *config.Settings, db database.Database, table *database.Table) (string, string) {

	var structFields strings.Builder

	columnInfo := columnInfo{}
	columns := map[string]struct{}{}

	for _, column := range table.Columns {

		columnName := strings.Title(column.Name)
		if settings.IsOutputFormatCamelCase() {
			columnName = camelCaseString(column.Name)
		}
		if settings.ShouldInitialism() {
			columnName = toInitialisms(columnName)
		}

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

		columnType, col := mapDbColumnTypeToGoType(settings, db, column)

		// save that we saw types of columns at least once
		if !columnInfo.isTemporal {
			columnInfo.isTemporal = col.isTemporal
		}
		if !columnInfo.isNullableTemporal {
			columnInfo.isNullableTemporal = col.isNullableTemporal
		}
		if !columnInfo.isNullablePrimitive {
			columnInfo.isNullablePrimitive = col.isNullablePrimitive
		}

		structFields.WriteString(columnName)
		structFields.WriteString(" ")
		structFields.WriteString(columnType)
		structFields.WriteString(" ")
		structFields.WriteString(taggers.GenerateTag(db, column))
		structFields.WriteString("\n")
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
	if settings.IsOutputFormatCamelCase() {
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

	if columnInfo.isNullablePrimitive && settings.IsNullTypeSQL() {
		content.WriteString("\t\"database/sql\"\n")
	}

	if columnInfo.isTemporal {
		content.WriteString("\t\"time\"\n")
	}

	if columnInfo.isNullableTemporal && settings.IsNullTypeSQL() {
		content.WriteString("\t\n")
		content.WriteString(db.GetDriverImportLibrary())
		content.WriteString("\n")
	}

	if settings.IsMastermindStructableRecorder {
		content.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
	}

	content.WriteString(")\n\n")
}

func mapDbColumnTypeToGoType(settings *config.Settings, db database.Database, column database.Column) (goType string, columnInfo columnInfo) {
	if db.IsString(column) || db.IsText(column) {
		goType = "string"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*string", "sql.NullString")
			columnInfo.isNullable = true
		}
	} else if db.IsInteger(column) {
		goType = "int"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*int", "sql.NullInt64")
			columnInfo.isNullable = true
		}
	} else if db.IsFloat(column) {
		goType = "float64"
		if db.IsNullable(column) {
			goType = getNullType(settings, "*float64", "sql.NullFloat64")
			columnInfo.isNullable = true
		}
	} else if db.IsTemporal(column) {
		if !db.IsNullable(column) {
			goType = "time.Time"
			columnInfo.isTemporal = true
		} else {
			goType = getNullType(settings, "*time.Time", db.GetTemporalDriverDataType())
			columnInfo.isTemporal = settings.Null == config.NullTypeNative
			columnInfo.isNullableTemporal = true
			columnInfo.isNullable = true
		}
	} else {
		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if db.IsNullable(column) {
				goType = getNullType(settings, "*bool", "sql.NullBool")
				columnInfo.isNullable = true
			}
		default:
			goType = getNullType(settings, "*string", "sql.NullString")
		}
	}

	columnInfo.isNullablePrimitive = columnInfo.isNullable && !db.IsTemporal(column)

	return goType, columnInfo
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
