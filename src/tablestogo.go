package tablestogo

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"bytes"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// a table has a name and a set (slice) of columns
type Table struct {
	TableName string `db:"table_name"`
	Columns   []Column
}

// stores information about a column
type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	ColumnName             string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	ColumnDefault          sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	ColumnKey              string         `db:"column_key"` // mysql specific
	Extra                  string         `db:"extra"`      // mysql specific
}

type Tagger interface {
	GenerateTag(column Column) string
}

type DbTag string
type StblTag string
type SqlTag string

func (t *DbTag) GenerateTag(column Column) string {
	return `db:"` + column.ColumnName + `"`
}

func (t *StblTag) GenerateTag(column Column) string {

	isPk := ""
	if strings.Contains(column.ColumnDefault.String, "nextval") || // pg
		(strings.Contains(column.ColumnKey, "PRI") && strings.Contains(column.Extra, "auto_increment")) { //mysql
		isPk = `,PRIMARY_KEY,SERIAL,AUTO_INCREMENT`
	}

	return `stbl:"` + column.ColumnName + isPk + `"`
}

// TODO
func (t *SqlTag) GenerateTag(column Column) string {
	return `sql:"` + column.ColumnName + `"`
}

var (
	// holds the db instance
	db *sqlx.DB

	SupportedDbTypes       = []string{"pg", "mysql"}
	SupportedOutputFormats = []string{"c", "o"}

	DbTypeToDriverMap = map[string]string{
		"pg":    "postgres",
		"mysql": "mysql",
	}

	DbDefaultPorts = map[string]string{
		"pg":    "5432",
		"mysql": "3306",
	}

	settings *Settings

	taggers = map[uint64]Tagger{
		1: new(DbTag),
		2: new(StblTag),
		4: new(SqlTag),
	}
)

// stores the supported settings / command line arguments
type Settings struct {
	Verbose        bool
	DbType         string
	User           string
	Pswd           string
	DbName         string
	Schema         string
	Host           string
	Port           string
	OutputFilePath string
	OutputFormat   string
	PackageName    string
	Prefix         string
	Suffix         string

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool

	TagsGorm bool

	effectiveTags uint64
}

// constructor for settings with default values
func NewSettings() *Settings {
	return &Settings{
		Verbose:        false,
		DbType:         "pg",
		User:           "postgres",
		Pswd:           "",
		DbName:         "postgres",
		Schema:         "public",
		Host:           "127.0.0.1",
		Port:           "5432",
		OutputFilePath: "./output",
		OutputFormat:   "c",
		PackageName:    "dto",
		Prefix:         "",
		Suffix:         "",

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		TagsGorm: false,

		effectiveTags: 1,
	}
}

// main function to run the conversions
func Run(s *Settings) (err error) {

	err = VerifySettings(s)
	if err != nil {
		return err
	}
	settings = s

	createEffectiveTags()

	err = connect()
	if err != nil {
		return err
	}
	defer db.Close()

	var database Database
	generalDatabase := &GeneralDatabase{
		db:       db,
		Settings: s,
	}

	switch s.DbType {
	case "mysql":
		database = &MySQLDatabase{
			GeneralDatabase: generalDatabase,
		}
	default: // pg
		database = &PostgreDatabase{
			GeneralDatabase: generalDatabase,
		}
	}

	return run(database)
}

// verifies the settings and checks the given output paths
func VerifySettings(settings *Settings) (err error) {

	if !IsStringInSlice(settings.DbType, SupportedDbTypes) {
		return errors.New(fmt.Sprintf("type of database %q not supported! %v", settings.DbType, SupportedDbTypes))
	}

	if !IsStringInSlice(settings.OutputFormat, SupportedOutputFormats) {
		return errors.New(fmt.Sprintf("output format %q not supported! %v", settings.OutputFormat, SupportedOutputFormats))
	}

	if err = verifyOutputPath(settings.OutputFilePath); err != nil {
		return err
	}

	if settings.OutputFilePath, err = prepareOutputPath(settings.OutputFilePath); err != nil {
		return err
	}

	if settings.Port == "" {
		settings.Port = DbDefaultPorts[settings.DbType]
	}

	if settings.PackageName == "" {
		return errors.New("name of package can not be empty!")
	}

	return err
}

func verifyOutputPath(outputFilePath string) (err error) {

	info, err := os.Stat(outputFilePath)

	if os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("output file path %q does not exists!", outputFilePath))
	}

	if !info.Mode().IsDir() {
		return errors.New(fmt.Sprintf("output file path %q is not a directory!", outputFilePath))
	}

	return err
}

func prepareOutputPath(ofp string) (outputFilePath string, err error) {
	outputFilePath, err = filepath.Abs(ofp + "/")
	return outputFilePath, err
}

func createEffectiveTags() {
	if settings.TagsNoDb {
		settings.effectiveTags = 0
	}
	if settings.TagsMastermindStructable {
		settings.effectiveTags |= 2
	}
	if settings.TagsMastermindStructableOnly {
		if !settings.TagsNoDb {
			settings.effectiveTags -= 1
		}
		if settings.TagsMastermindStructable {
			settings.effectiveTags -= 2
		}
		settings.effectiveTags |= 4
	}
	if settings.TagsGorm {
		settings.effectiveTags |= 8
	}
}

func connect() (err error) {
	db, err = sqlx.Connect(DbTypeToDriverMap[settings.DbType], createDataSourceName())
	if err != nil {
		usingPswd := "no"
		if settings.Pswd != "" {
			usingPswd = "yes"
		}
		return errors.New(
			fmt.Sprintf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed:\r\n%v",
				settings.DbType, settings.User, settings.DbName, settings.Host, settings.Port, usingPswd, err))
	}
	return db.Ping()
}

func createDataSourceName() (dataSourceName string) {
	switch settings.DbType {
	case "mysql":
		dataSourceName = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", settings.User, settings.Pswd, settings.Host, settings.Port, settings.DbName)
	default: // pg
		dataSourceName = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			settings.Host, settings.Port, settings.User, settings.DbName, settings.Pswd)
	}
	return dataSourceName
}

func run(db Database) (err error) {

	fmt.Printf("running for %q...\r\n", settings.DbType)

	tables, err := db.GetTables()

	if err != nil {
		return err
	}

	if settings.Verbose {
		fmt.Printf("> count of tables: %v\r\n", len(tables))
	}

	err = db.PrepareGetColumnsOfTableStmt()

	if err != nil {
		return err
	}

	for _, table := range tables {

		if settings.Verbose {
			fmt.Printf("> processing table %q\r\n", table.TableName)
		}

		err = db.GetColumnsOfTable(table)

		if err != nil {
			return err
		}

		err = createStructOfTable(table)

		if err != nil {
			if settings.Verbose {
				fmt.Printf(">Error at createStructOfTable(%v)\r\n", table.TableName)
			}
			return err
		}
	}

	fmt.Println("done!")

	return err
}

func createStructOfTable(table *Table) (err error) {

	var fileContentBuffer, structFieldsBuffer bytes.Buffer
	var isNullable bool
	timeIndicator := 0

	for _, column := range table.Columns {

		columnName := strings.Title(column.ColumnName)
		if settings.OutputFormat == "c" {
			columnName = CamelCaseString(columnName)
		}
		columnType, isTime := mapDbColumnTypeToGoType(column.DataType, column.IsNullable)

		structFieldsBuffer.WriteString("\t" + columnName + " " + columnType + generateTags(column) + "\n")

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
		tableName = CamelCaseString(tableName)
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

func generateTags(column Column) (tags string) {
	var t uint64
	for t = 1; t <= settings.effectiveTags; t *= 2 {
		if shouldTag(t) {
			tags += taggers[t].GenerateTag(column) + " "
		}
	}
	if len(tags) > 0 {
		tags = " `" + strings.TrimSpace(tags) + "`"
	}
	return tags
}

func shouldTag(t uint64) bool {
	return settings.effectiveTags&t > 0
}

func mapDbColumnTypeToGoType(dbDataType string, isNullable string) (goType string, isTime bool) {

	isTime = false

	// first row: postgresql datatypes  // TODO bitstrings, enum, other special types
	// second row: additional mysql datatypes not covered by first row // TODO bit, enums, set
	// and so on

	switch dbDataType {
	case "integer", "bigint", "bigserial", "smallint", "smallserial", "serial",
		"int", "tinyint", "mediumint":
		goType = "int"
		if isNullable == "YES" {
			goType = "sql.NullInt64"
		}
	case "double precision", "numeric", "decimal", "real",
		"float", "double":
		goType = "float64"
		if isNullable == "YES" {
			goType = "sql.NullFloat64"
		}
	case "character varying", "character", "text",
		"char", "varchar", "binary", "varbinary", "blob":
		goType = "string"
		if isNullable == "YES" {
			goType = "sql.NullString"
		}
	case "time", "timestamp", "time with time zone", "timestamp with time zone", "time without time zone", "timestamp without time zone",
		"date", "datetime", "year":
		goType = "time.Time"
		if isNullable == "YES" {
			goType = "pq.NullTime"
		}
		isTime = true
	case "boolean":
		goType = "bool"
		if isNullable == "YES" {
			goType = "sql.NullBool"
		}
	default:
		goType = "sql.NullString"
	}

	return goType, isTime
}
