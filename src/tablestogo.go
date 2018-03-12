package tablestogo

import (
	"bytes"
	"database/sql"
	"fmt"
	"go/format"
	"os"
	"strings"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	// postgres database driver
	_ "github.com/lib/pq"
)

var (
	// dbTypeToDriverMap maps the database type to the driver names
	dbTypeToDriverMap = map[string]string{
		"pg":    "postgres",
		"mysql": "mysql",
	}

	// map of Tagger used
	// key is a ascending sequence of i*2 to determine easily which tags to generate later
	taggers = map[int]Tagger{
		1: new(DbTag),
		2: new(StblTag),
		4: new(SQLTag),
	}
)

// Table has a name and a set (slice) of columns
type Table struct {
	TableName string `db:"table_name"`
	Columns   []Column
}

// Column stores information about a column
type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	ColumnName             string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	ColumnDefault          sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	DatetimePrecision      sql.NullInt64  `db:"datetime_precision"`
	ColumnKey              string         `db:"column_key"`      // mysql specific
	Extra                  string         `db:"extra"`           // mysql specific
	ConstraintName         sql.NullString `db:"constraint_name"` // pg specific
	ConstraintType         sql.NullString `db:"constraint_type"` // pg specific
}

// Tagger interface for types of struct-tages
type Tagger interface {
	GenerateTag(db Database, column Column) string
}

// DbTag is the standard "db"-tag
type DbTag string

// GenerateTag for DbTag to satisfy the Tagger interface
func (t *DbTag) GenerateTag(db Database, column Column) string {
	return `db:"` + column.ColumnName + `"`
}

// StblTag represents the Masterminds/structable "stbl"-tag
type StblTag string

// GenerateTag for StblTag to satisfy the Tagger interface
func (t *StblTag) GenerateTag(db Database, column Column) string {

	isPk := ""
	if db.IsPrimaryKey(column) {
		isPk = ",PRIMARY_KEY"
	}

	isAutoIncrement := ""
	if db.IsAutoIncrement(column) {
		isAutoIncrement = ",SERIAL,AUTO_INCREMENT"
	}

	return `stbl:"` + column.ColumnName + isPk + isAutoIncrement + `"`
}

// SQLTag is the experimental "sql"-tag
type SQLTag string

// GenerateTag for SQLTag to satisfy the Tagger interface
func (t *SQLTag) GenerateTag(db Database, column Column) string {

	colType := ""
	characterMaximumLength := ""
	if db.IsString(column) && column.CharacterMaximumLength.Valid {
		characterMaximumLength = fmt.Sprintf("(%v)", column.CharacterMaximumLength.Int64)
	}

	colType = fmt.Sprintf("type:%v%v;", column.DataType, characterMaximumLength)

	isNullable := ""
	if !db.IsNullable(column) {
		isNullable = "not null;"
	}

	// TODO size:###
	// TODO unique, key, index, ...

	tag := colType + isNullable
	tag = strings.TrimSuffix(tag, ";")

	return `sql:"` + tag + `"`
}

// Run is the main function to run the conversions
func Run(settings *Settings) (err error) {

	createEffectiveTags(settings)

	database := makeDatabase(settings)

	err = connect(settings, database)
	if err != nil {
		return err
	}
	defer database.Db.Close()

	return run(database)
}

// factory function
func makeDatabase(settings *Settings) *GeneralDatabase {

	database := &GeneralDatabase{
		settings: settings,
	}

	// FIXME ugly!?!?!?! cyclic dependency !!??
	switch settings.DbType {
	case "mysql":
		mysql := &MySQLDatabase{}
		database.Database = mysql
		mysql.GeneralDatabase = database
	default: // pg
		pg := &PostgreDatabase{}
		database.Database = pg
		pg.GeneralDatabase = database
	}

	return database
}

func createEffectiveTags(settings *Settings) {
	if settings.TagsNoDb {
		settings.effectiveTags = 0
	}
	if settings.TagsMastermindStructable {
		settings.effectiveTags |= 2
	}
	if settings.TagsMastermindStructableOnly {
		settings.effectiveTags = 0
		settings.effectiveTags |= 2
	}
	if settings.TagsSQL {
		settings.effectiveTags |= 4
	}
	if settings.TagsSQLOnly {
		settings.effectiveTags = 0
		settings.effectiveTags |= 4
	}
	// last tag-"ONLY" wins if multiple specified
}

func connect(settings *Settings, database *GeneralDatabase) (err error) {
	database.Db, err = sqlx.Connect(dbTypeToDriverMap[settings.DbType], database.CreateDataSourceName(settings))
	if err != nil {
		usingPswd := "no"
		if settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed:\r\n%v",
			settings.DbType, settings.User, settings.DbName, settings.Host, settings.Port, usingPswd, err)
	}
	return database.Db.Ping()
}

func run(database *GeneralDatabase) (err error) {

	fmt.Printf("running for %q...\r\n", database.settings.DbType)

	tables, err := database.GetTables()

	if err != nil {
		return err
	}

	if database.settings.Verbose {
		fmt.Printf("> number of tables: %v\r\n", len(tables))
	}

	err = database.PrepareGetColumnsOfTableStmt()

	if err != nil {
		return err
	}

	for _, table := range tables {

		if database.settings.Verbose {
			fmt.Printf("> processing table %q\r\n", table.TableName)
		}

		err = database.GetColumnsOfTable(table)

		if err != nil {
			return err
		}

		if database.settings.Verbose {
			fmt.Printf("\t> number of columns: %v\r\n", len(table.Columns))
		}

		err = createStructOfTable(database, table)

		if err != nil {
			if database.settings.Verbose {
				fmt.Printf(">Error at createStructOfTable(%v)\r\n", table.TableName)
			}
			return err
		}
	}

	fmt.Println("done!")

	return err
}

func createStructOfTable(database *GeneralDatabase, table *Table) (err error) {

	var fileContentBuffer, structFieldsBuffer bytes.Buffer
	var isNullable bool
	timeIndicator := 0

	for _, column := range table.Columns {

		// TODO add verbosity levels
		//if settings.Verbose {
		//	fmt.Printf("\t> %v\r\n", column.ColumnName)
		//}

		columnName := strings.Title(column.ColumnName)
		if database.settings.OutputFormat == "c" {
			columnName = CamelCaseString(columnName)
		}
		columnType, isTime := mapDbColumnTypeToGoType(database, column)

		structFieldsBuffer.WriteString("\t" + columnName + " " + columnType + generateTags(database, column) + "\n")

		// collect some info for later use
		if column.IsNullable == "YES" {
			isNullable = true
		}
		if isTime {
			timeIndicator++
		}
	}

	if database.settings.IsMastermindStructableRecorder {
		structFieldsBuffer.WriteString("\t\nstructable.Recorder\n")
	}

	// create file
	tableName := strings.Title(database.settings.Prefix + table.TableName + database.settings.Suffix)
	if database.settings.OutputFormat == "c" {
		tableName = CamelCaseString(tableName)
	}

	outFile, err := os.Create(database.settings.OutputFilePath + tableName + ".go")

	if err != nil {
		return err
	}

	// write header infos
	fileContentBuffer.WriteString("package " + database.settings.PackageName + "\n\n")

	// do imports
	if isNullable || timeIndicator > 0 || database.settings.IsMastermindStructableRecorder {
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

		if database.settings.IsMastermindStructableRecorder {
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

func generateTags(database *GeneralDatabase, column Column) (tags string) {
	for t := 1; t <= database.settings.effectiveTags; t *= 2 {
		if shouldTag(database, t) {
			tags += taggers[t].GenerateTag(database, column) + " "
		}
	}
	if len(tags) > 0 {
		tags = " `" + strings.TrimSpace(tags) + "`"
	}
	return tags
}

func shouldTag(database *GeneralDatabase, t int) bool {
	return database.settings.effectiveTags&t > 0
}

func mapDbColumnTypeToGoType(database Database, column Column) (goType string, isTime bool) {

	isTime = false

	if database.IsString(column) || database.IsText(column) {
		goType = "string"
		if database.IsNullable(column) {
			goType = "sql.NullString"
		}
	} else if database.IsInteger(column) {
		goType = "int"
		if database.IsNullable(column) {
			goType = "sql.NullInt64"
		}
	} else if database.IsFloat(column) {
		goType = "float64"
		if database.IsNullable(column) {
			goType = "sql.NullFloat64"
		}
	} else if database.IsTemporal(column) {
		goType = "time.Time"
		if database.IsNullable(column) {
			goType = "pq.NullTime"
		}
		isTime = true
	} else {

		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if database.IsNullable(column) {
				goType = "sql.NullBool"
			}
		default:
			goType = "sql.NullString"
		}
	}

	return goType, isTime
}
