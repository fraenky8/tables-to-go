package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/output"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
	"github.com/fraenky8/tables-to-go/v2/pkg/tagger"
)

var (
	// some strings for idiomatic go in column names
	// see https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	initialisms = []string{"ID", "JSON", "XML", "HTTP", "URL", "UUID"}
)

// App is the dependency container for this CLI tool.
type App struct {
	settings *settings.Settings
	taggers  *tagger.Taggers
	db       database.Database
	out      output.Writer
	caser    cases.Caser
}

// New creates a new App.
func New(s *settings.Settings, db database.Database, out output.Writer) *App {
	return &App{
		settings: s,
		taggers:  tagger.NewTaggers(s),
		db:       db,
		out:      out,
		caser:    cases.Title(language.English, cases.NoLower),
	}
}

// Run runs the transformations by creating the concrete Database by the provided settings
func (app *App) Run(ctx context.Context) error {
	app.printf("running for %q...\r\n", app.settings.DbType)

	tables, err := app.db.GetTables(ctx, app.settings.Tables...)
	if err != nil {
		return fmt.Errorf("could not get tables: %w", err)
	}

	if app.settings.Verbose {
		app.printf("> number of tables: %v\r\n", len(tables))
	}

	if err = app.db.PrepareGetColumnsOfTableStmt(ctx); err != nil {
		return fmt.Errorf("could not prepare the get-column-statement: %w", err)
	}

	for _, table := range tables {
		select {
		case <-ctx.Done():
			if app.settings.Verbose {
				app.printf("> received cancellation: %v\r\n", context.Cause(ctx))
			}
			return ctx.Err()
		default:
		}

		if app.settings.Verbose {
			app.printf("> processing table %q\r\n", table.Name)
		}

		if err = app.db.GetColumnsOfTable(ctx, table); err != nil {
			if !app.settings.Force {
				return fmt.Errorf("could not get columns of table %q: %w", table.Name, err)
			}
			app.printf("could not get columns of table %q: %v\n", table.Name, err)
			continue
		}

		if app.settings.Verbose {
			app.printf("\t> number of columns: %v\r\n", len(table.Columns))
		}

		tableName, content, err := app.createTableStructString(table)

		if err != nil {
			if !app.settings.Force {
				return fmt.Errorf("could not create string for table %q: %w", table.Name, err)
			}
			app.printf("could not create string for table %q: %v\n", table.Name, err)
			continue
		}

		fileName := app.camelCaseString(tableName)
		if app.settings.IsFileNameFormatSnakeCase() {
			fileName = strcase.ToSnake(fileName)
		}

		err = app.out.Write(fileName, content)
		if err != nil {
			if !app.settings.Force {
				return fmt.Errorf("could not write struct for table %q: %w", table.Name, err)
			}
			app.printf("could not write struct for table %q: %v\n", table.Name, err)
		}
	}

	app.println("done!")

	return nil
}

type columnInfo struct {
	isNullable bool
	isTemporal bool
}

func (c columnInfo) isNullableOrTemporal() bool {
	return c.isNullable || c.isTemporal
}

func (app *App) createTableStructString(table *database.Table) (string, string, error) {

	tableName := app.caser.String(app.settings.Prefix) + app.caser.String(table.Name) + app.caser.String(app.settings.Suffix)
	// Replace any whitespace with underscores
	tableName = strings.Map(replaceSpace, tableName)
	if app.settings.IsOutputFormatCamelCase() {
		tableName = app.camelCaseString(tableName)
	}

	// Check that the table name doesn't contain any invalid characters for Go variables
	if !validVariableName(tableName) {
		return "", "", fmt.Errorf("table name %q contains invalid characters", table.Name)
	}

	var (
		structFields strings.Builder
		columnInfo   columnInfo
		columns      = make(map[string]struct{}, len(table.Columns))
	)
	for _, column := range table.Columns {
		columnName, err := app.formatColumnName(column.Name, table.Name)
		if err != nil {
			return "", "", err
		}

		// ISSUE-4: if columns are part of multiple constraints
		// then the sql returns multiple rows per column name.
		// Therefore, we check if we already added a column with
		// that name to the struct, if so, skip.
		if _, ok := columns[columnName]; ok {
			continue
		}
		columns[columnName] = struct{}{}

		if app.settings.VVerbose {
			app.printf("\t\t> %v\r\n", column.Name)
		}

		columnType, col := app.mapDbColumnTypeToGoType(column)

		// save that we saw types of columns at least once
		if !columnInfo.isTemporal {
			columnInfo.isTemporal = col.isTemporal
		}
		if !columnInfo.isNullable {
			columnInfo.isNullable = col.isNullable
		}

		structFields.WriteString(columnName)
		structFields.WriteString(" ")
		structFields.WriteString(columnType)
		structFields.WriteString(" ")
		structFields.WriteString(app.taggers.GenerateTag(app.db, column))
		structFields.WriteString("\n")
	}

	if app.settings.IsMastermindStructableRecorder {
		structFields.WriteString("\t\nstructable.Recorder\n")
	}

	var fileContent strings.Builder

	// write header infos
	fileContent.WriteString("package ")
	fileContent.WriteString(app.settings.PackageName)
	fileContent.WriteString("\n\n")

	// write imports
	app.generateImports(&fileContent, columnInfo)

	// write struct with fields
	fileContent.WriteString("type ")
	fileContent.WriteString(tableName)
	fileContent.WriteString(" struct {\n")
	fileContent.WriteString(structFields.String())
	fileContent.WriteString("}")

	return tableName, fileContent.String(), nil
}

func (app *App) generateImports(content *strings.Builder, columnInfo columnInfo) {

	if !columnInfo.isNullableOrTemporal() && !app.settings.IsMastermindStructableRecorder {
		return
	}

	content.WriteString("import (\n")

	if columnInfo.isNullable && app.settings.IsNullTypeSQL() {
		content.WriteString("\t\"database/sql\"\n")
	}

	if columnInfo.isTemporal {
		content.WriteString("\t\"time\"\n")
	}

	if app.settings.IsMastermindStructableRecorder {
		content.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
	}

	content.WriteString(")\n\n")
}

func (app *App) mapDbColumnTypeToGoType(column database.Column) (goType string, columnInfo columnInfo) {
	if app.db.IsInteger(column) {
		goType = "int"
		if app.db.IsNullable(column) {
			goType = getNullType(app.settings, "*int", "sql.NullInt64")
			columnInfo.isNullable = true
		}
	} else if app.db.IsFloat(column) {
		goType = "float64"
		if app.db.IsNullable(column) {
			goType = getNullType(app.settings, "*float64", "sql.NullFloat64")
			columnInfo.isNullable = true
		}
	} else if app.db.IsTemporal(column) {
		if app.db.IsNullable(column) {
			goType = getNullType(app.settings, "*time.Time", "sql.NullTime")
			columnInfo.isTemporal = !app.settings.IsNullTypeSQL()
			columnInfo.isNullable = true
		} else {
			goType = "time.Time"
			columnInfo.isTemporal = true
		}
	} else {
		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if app.db.IsNullable(column) {
				goType = getNullType(app.settings, "*bool", "sql.NullBool")
				columnInfo.isNullable = true
			}
		default:
			// Everything else we cannot detect defaults to (nullable) string.
			goType = "string"
			if app.db.IsNullable(column) {
				goType = getNullType(app.settings, "*string", "sql.NullString")
				columnInfo.isNullable = true
			}
		}
	}

	return goType, columnInfo
}

func (app *App) camelCaseString(s string) string {
	if s == "" {
		return s
	}

	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		return app.caser.String(s)
	}

	var cc strings.Builder
	for _, part := range splitted {
		cc.WriteString(app.caser.String(strings.ToLower(part)))
	}
	return cc.String()
}

func getNullType(settings *settings.Settings, primitive, sql string) string {
	if settings.IsNullTypeSQL() {
		return sql
	}
	return primitive
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
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
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

// FormatColumnName checks for invalid characters and transforms a column name
// according to the provided settings.
func (app *App) formatColumnName(column, table string) (string, error) {

	// Replace any whitespace with underscores
	columnName := strings.Map(replaceSpace, column)
	columnName = app.caser.String(columnName)

	if app.settings.IsOutputFormatCamelCase() {
		columnName = app.camelCaseString(columnName)
	}
	if app.settings.ShouldInitialism() {
		columnName = toInitialisms(columnName)
	}

	// Check that the column name doesn't contain any invalid characters for Go variables
	if !validVariableName(columnName) {
		return "", fmt.Errorf("column name %q in table %q contains invalid characters", column, table)
	}

	// First character of an identifier in Go must be letter or _
	// We want it to be an uppercase letter to be a public field
	if !unicode.IsLetter(rune(columnName[0])) {
		prefix := "X_"
		if app.settings.IsOutputFormatCamelCase() {
			prefix = "X"
		}
		if app.settings.ShouldInitialism() {
			// Note we use the original passed in name of the column here to
			// avoid the Title'izing of the first non-digit character as done
			// by cases.Caser. Eg: `1fish2fish` gets transformed to `X1Fish2fish`
			// but we want `X1fish2fish`.
			columnName = toInitialisms(column)
		}
		if app.settings.Verbose {
			app.printf("\t\t>column %q in table %q doesn't start with a letter; prepending with %q\n", column, table, prefix)
		}
		columnName = prefix + columnName
	}

	return columnName, nil
}

func (app *App) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}

func (app *App) println(a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}
