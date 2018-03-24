package tablestogo

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/fraenky8/tables-to-go/src/database"
	"github.com/fraenky8/tables-to-go/src/settings"
	"github.com/fraenky8/tables-to-go/src/tagger"
)

var (
	// map of Tagger used
	// key is a ascending sequence of i*2 to determine easily which tags to generate later
	taggers = map[int]tagger.Tagger{
		1: new(tagger.DbTag),
		2: new(tagger.StblTag),
		4: new(tagger.SQLTag),
	}

	effectiveTags = 1
)

// Run is the main function to run the conversions
func Run(settings *settings.Settings) (err error) {

	createEffectiveTags(settings)

	db := database.NewDatabase(settings)

	if err = db.Connect(); err != nil {
		return err
	}
	defer db.Close()

	return run(settings, db)
}

func run(settings *settings.Settings, db database.Database) (err error) {

	fmt.Printf("running for %q...\r\n", settings.DbType)

	tables, err := db.GetTables()
	if err != nil {
		return err
	}

	if settings.Verbose {
		fmt.Printf("> number of tables: %v\r\n", len(tables))
	}

	if err = db.PrepareGetColumnsOfTableStmt(); err != nil {
		return err
	}

	for _, table := range tables {

		if settings.Verbose {
			fmt.Printf("> processing table %q\r\n", table.TableName)
		}

		if err = db.GetColumnsOfTable(table); err != nil {
			return err
		}

		if settings.Verbose {
			fmt.Printf("\t> number of columns: %v\r\n", len(table.Columns))
		}

		if err = createStructOfTable(settings, db, table); err != nil {
			if settings.Verbose {
				fmt.Printf(">Error at createStructOfTable(%v)\r\n", table.TableName)
			}
			return err
		}
	}

	fmt.Println("done!")

	return err
}

func createStructOfTable(settings *settings.Settings, db database.Database, table *database.Table) (err error) {

	var fileContentBuffer, structFieldsBuffer bytes.Buffer
	var isNullable bool
	timeIndicator := 0

	for _, column := range table.Columns {

		// TODO add verbosity levels
		// if settings.Verbose {
		// 	fmt.Printf("\t> %v\r\n", column.ColumnName)
		// }

		columnName := strings.Title(column.ColumnName)
		if settings.OutputFormat == "c" {
			columnName = camelCaseString(columnName)
		}
		columnType, isTime := mapDbColumnTypeToGoType(db, column)

		structFieldsBuffer.WriteString("\t" + columnName + " " + columnType + generateTags(db, column) + "\n")

		// collect some info for later use
		if column.IsNullable == "YES" {
			isNullable = true
		}
		if isTime {
			timeIndicator++
		}
	}

	if settings.IsMastermindStructableRecorder {
		structFieldsBuffer.WriteString("\t\nstructable.Recorder\n")
	}

	// create file
	tableName := strings.Title(settings.Prefix + table.TableName + settings.Suffix)
	if settings.OutputFormat == "c" {
		tableName = camelCaseString(tableName)
	}

	outFile, err := os.Create(settings.OutputFilePath + tableName + ".go")

	if err != nil {
		return err
	}

	// write header infos
	fileContentBuffer.WriteString("package " + settings.PackageName + "\n\n")

	// do imports
	if isNullable || timeIndicator > 0 || settings.IsMastermindStructableRecorder {
		fileContentBuffer.WriteString("import (\n")

		if isNullable {
			fileContentBuffer.WriteString("\t\"database/sql\"\n")
		}

		if timeIndicator > 0 {
			if isNullable {
				fileContentBuffer.WriteString("\t\n\"github.com/lib/pq\"\n")
			} else {
				fileContentBuffer.WriteString("\t\"time\"\n")
			}
		}

		if settings.IsMastermindStructableRecorder {
			fileContentBuffer.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
		}

		fileContentBuffer.WriteString(")\n\n")
	}

	// write struct with fields
	fileContentBuffer.WriteString("type " + tableName + " struct {\n")
	fileContentBuffer.WriteString(structFieldsBuffer.String())
	fileContentBuffer.WriteString("}")

	// format it
	formatedFile, _ := format.Source(fileContentBuffer.Bytes())

	// and save it in file
	outFile.Write(formatedFile)
	outFile.Sync()
	outFile.Close()

	return err
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
