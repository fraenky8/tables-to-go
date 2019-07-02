package cli

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
	"github.com/fraenky8/tables-to-go/pkg/settings"
	"github.com/fraenky8/tables-to-go/pkg/tagger"
)

var (
	taggers tagger.Tagger

	// some strings for idiomatic go in column names
	// see https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	initialisms = []string{"ID", "JSON", "XML", "HTTP", "URL"}
)

// Run runs the transformations by creating the concrete Database by the provided settings
func Run(settings *settings.Settings, db database.Database, out output.Writer) (err error) {

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
			if !settings.Force {
				return fmt.Errorf("could not get columns of table %q: %v", table.Name, err)
			}
			fmt.Printf("could not get columns of table %q: %v\n", table.Name, err)
			continue
		}

		if settings.Verbose {
			fmt.Printf("\t> number of columns: %v\r\n", len(table.Columns))
		}

		tableName, content, err := createTableStructString(settings, db, table)
		if err != nil {
			if !settings.Force {
				return fmt.Errorf("could not create string for table %q: %v", table.Name, err)
			}
			fmt.Printf("could not create string for table %q: %v\n", table.Name, err)
			continue
		}

		err = out.Write(tableName, content)
		if err != nil {
			if !settings.Force {
				return fmt.Errorf("could not write struct for table %q: %v", table.Name, err)
			}
			fmt.Printf("could not write struct for table %q: %v\n", table.Name, err)
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

func createTableStructString(settings *settings.Settings, db database.Database, table *database.Table) (string, string, error) {

	var structFields strings.Builder
	tableName := strings.Title(settings.Prefix + table.Name + settings.Suffix)
	// Replace any whitespace with underscores
	tableName = strings.Map(replaceSpace, tableName)
	if settings.IsOutputFormatCamelCase() {
		tableName = camelCaseString(tableName)
	}

	// Check that the table name doesn't contain any invalid characters for Go variables
	if !validVariableName(tableName) {
		return "", "", fmt.Errorf("table name %q contains invalid characters", table.Name)
	}

	columnInfo := columnInfo{}
	columns := map[string]struct{}{}

	for _, column := range table.Columns {

		columnName := strings.Title(column.Name)
		// Replace any whitespace with underscores
		columnName = strings.Map(replaceSpace, columnName)
		if settings.IsOutputFormatCamelCase() {
			columnName = camelCaseString(columnName)
		}
		if settings.ShouldInitialism() {
			columnName = toInitialisms(columnName)
		}

		// Check that the column name doesn't contain any invalid characters for Go variables
		if !validVariableName(columnName) {
			return "", "", fmt.Errorf("column name %q in table %q contains invalid characters", column.Name, table.Name)
		}
		// First character of an identifier in Go must be letter or _
		// We want it to be an uppercase letter to be a public field
		if !unicode.IsLetter([]rune(columnName)[0]) {
			if settings.Verbose {
				fmt.Printf("\t\t>column %q in table %q doesn't start with a letter; prepending with X_\n", column.Name, table.Name)
			}
			columnName = "X_" + columnName
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

	// write struct with fields
	fileContent.WriteString("type ")
	fileContent.WriteString(tableName)
	fileContent.WriteString(" struct {\n")
	fileContent.WriteString(structFields.String())
	fileContent.WriteString("}")

	return tableName, fileContent.String(), nil
}

func generateImports(content *strings.Builder, settings *settings.Settings, db database.Database, columnInfo columnInfo) {

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

func mapDbColumnTypeToGoType(s *settings.Settings, db database.Database, column database.Column) (goType string, columnInfo columnInfo) {
	if db.IsString(column) || db.IsText(column) {
		goType = "string"
		if db.IsNullable(column) {
			goType = getNullType(s, "*string", "sql.NullString")
			columnInfo.isNullable = true
		}
	} else if db.IsInteger(column) {
		goType = "int"
		if db.IsNullable(column) {
			goType = getNullType(s, "*int", "sql.NullInt64")
			columnInfo.isNullable = true
		}
	} else if db.IsFloat(column) {
		goType = "float64"
		if db.IsNullable(column) {
			goType = getNullType(s, "*float64", "sql.NullFloat64")
			columnInfo.isNullable = true
		}
	} else if db.IsTemporal(column) {
		if !db.IsNullable(column) {
			goType = "time.Time"
			columnInfo.isTemporal = true
		} else {
			goType = getNullType(s, "*time.Time", db.GetTemporalDriverDataType())
			columnInfo.isTemporal = s.Null == settings.NullTypeNative
			columnInfo.isNullableTemporal = true
			columnInfo.isNullable = true
		}
	} else {
		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if db.IsNullable(column) {
				goType = getNullType(s, "*bool", "sql.NullBool")
				columnInfo.isNullable = true
			}
		default:
			goType = getNullType(s, "*string", "sql.NullString")
		}
	}

	columnInfo.isNullablePrimitive = columnInfo.isNullable && !db.IsTemporal(column)

	return goType, columnInfo
}

func getNullType(settings *settings.Settings, primitive string, sql string) string {
	if settings.IsNullTypeSQL() {
		return sql
	}
	return primitive
}

func camelCaseString(s string) string {
	if s == "" {
		return s
	}

	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		return strings.Title(s)
	}

	var cc string
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

// ValidVariableName checks for the existence of any characters
// outside of Unicode letters, numbers and underscore.
func validVariableName(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			return false
		}
	}
	return true
}

// ReplaceSpace swaps any Unicode space characters for underscores
// to create valid Go identifiers
func replaceSpace(r rune) rune {
	if unicode.IsSpace(r) || r == '\u200B' {
		return '_'
	}
	return r
}
