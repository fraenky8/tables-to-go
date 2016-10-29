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

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// postgres database driver
	_ "github.com/lib/pq"
)

var (
	// holds the db instance
	db *sqlx.DB

	// used concrete database, one of the supported types below
	database Database

	// the global applied settings
	settings *Settings

	// SupportedDbTypes represents the supported databases
	SupportedDbTypes = []string{"pg", "mysql"}
	// SupportedOutputFormats represents the supported output formats
	SupportedOutputFormats = []string{"c", "o"}

	// DbTypeToDriverMap maps the database type to the driver names
	DbTypeToDriverMap = map[string]string{
		"pg":    "postgres",
		"mysql": "mysql",
	}

	// DbDefaultPorts maps the database type to the default ports
	DbDefaultPorts = map[string]string{
		"pg":    "5432",
		"mysql": "3306",
	}

	// map of Tagger used
	// key is a ascending sequence of i*2 to determine easily which tags to generate later
	taggers = map[int]Tagger{
		1: new(DbTag),
		2: new(StblTag),
		4: new(SQLTag),
	}
)

// Settings stores the supported settings / command line arguments
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

	// TODO not implemented yet
	TagsGorm bool

	// experimental
	TagsSQL     bool
	TagsSQLOnly bool

	effectiveTags int
}

// NewSettings constructs settings with default values
func NewSettings() *Settings {
	return &Settings{
		Verbose:        false,
		DbType:         "pg",
		User:           "postgres",
		Pswd:           "",
		DbName:         "postgres",
		Schema:         "public",
		Host:           "127.0.0.1",
		Port:           "", // left blank -> is automatically determined if not set
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

		TagsSQL:     false,
		TagsSQLOnly: false,

		effectiveTags: 1,
	}
}

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
	GenerateTag(column Column) string
}

// DbTag is the standard "db"-tag
type DbTag string

// GenerateTag for DbTag to satisfy the Tagger interface
func (t *DbTag) GenerateTag(column Column) string {
	return `db:"` + column.ColumnName + `"`
}

// StblTag represents the Masterminds/structable "stbl"-tag
type StblTag string

// GenerateTag for StblTag to satisfy the Tagger interface
func (t *StblTag) GenerateTag(column Column) string {

	isPk := ""
	if database.IsPrimaryKey(column) {
		isPk = ",PRIMARY_KEY"
	}

	isAutoIncrement := ""
	if database.IsAutoIncrement(column) {
		isAutoIncrement = ",SERIAL,AUTO_INCREMENT"
	}

	return `stbl:"` + column.ColumnName + isPk + isAutoIncrement + `"`
}

// SQLTag is the experimental "sql"-tag
type SQLTag string

// GenerateTag for SQLTag to satisfy the Tagger interface
func (t *SQLTag) GenerateTag(column Column) string {

	colType := ""
	characterMaximumLength := ""
	if database.IsString(column) && column.CharacterMaximumLength.Valid {
		characterMaximumLength = fmt.Sprintf("(%v)", column.CharacterMaximumLength.Int64)
	}

	colType = fmt.Sprintf("type:%v%v;", column.DataType, characterMaximumLength)

	isNullable := ""
	if !database.IsNullable(column) {
		isNullable = "not null;"
	}

	// TODO size:###
	// TODO unique, key, index, ...

	tag := colType + isNullable
	tag = strings.TrimSuffix(tag, ";")

	return `sql:"` + tag + `"`
}

// Run is the main function to run the conversions
func Run(s *Settings) (err error) {

	err = VerifySettings(s)
	if err != nil {
		return err
	}
	settings = s

	createEffectiveTags()

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

	// connection must be appear here, database must exists at this point
	err = connect()
	if err != nil {
		return err
	}
	defer db.Close()

	return run()
}

// VerifySettings verifies the settings and checks the given output paths
func VerifySettings(settings *Settings) (err error) {

	if !IsStringInSlice(settings.DbType, SupportedDbTypes) {
		return fmt.Errorf("type of database %q not supported! %v", settings.DbType, SupportedDbTypes)
	}

	if !IsStringInSlice(settings.OutputFormat, SupportedOutputFormats) {
		return fmt.Errorf("output format %q not supported! %v", settings.OutputFormat, SupportedOutputFormats)
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
		return errors.New("name of package can not be empty")
	}

	return err
}

func verifyOutputPath(outputFilePath string) (err error) {

	info, err := os.Stat(outputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("output file path %q does not exists", outputFilePath)
	}

	if !info.Mode().IsDir() {
		return fmt.Errorf("output file path %q is not a directory", outputFilePath)
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

func connect() (err error) {
	db, err = sqlx.Connect(DbTypeToDriverMap[settings.DbType], database.CreateDataSourceName(settings))
	if err != nil {
		usingPswd := "no"
		if settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed:\r\n%v",
			settings.DbType, settings.User, settings.DbName, settings.Host, settings.Port, usingPswd, err)
	}
	return db.Ping()
}

func run() (err error) {

	fmt.Printf("running for %q...\r\n", settings.DbType)

	tables, err := database.GetTables()

	if err != nil {
		return err
	}

	if settings.Verbose {
		fmt.Printf("> count of tables: %v\r\n", len(tables))
	}

	err = database.PrepareGetColumnsOfTableStmt()

	if err != nil {
		return err
	}

	for _, table := range tables {

		if settings.Verbose {
			fmt.Printf("> processing table %q\r\n", table.TableName)
		}

		err = database.GetColumnsOfTable(table)

		if err != nil {
			return err
		}

		if settings.Verbose {
			fmt.Printf("\t> count of columns: %v\r\n", len(table.Columns))
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

		// TODO add verbosity levels
		//if settings.Verbose {
		//	fmt.Printf("\t> %v\r\n", column.ColumnName)
		//}

		columnName := strings.Title(column.ColumnName)
		if settings.OutputFormat == "c" {
			columnName = CamelCaseString(columnName)
		}
		columnType, isTime := mapDbColumnTypeToGoType(column)

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
	for t := 1; t <= settings.effectiveTags; t *= 2 {
		if shouldTag(t) {
			tags += taggers[t].GenerateTag(column) + " "
		}
	}
	if len(tags) > 0 {
		tags = " `" + strings.TrimSpace(tags) + "`"
	}
	return tags
}

func shouldTag(t int) bool {
	return settings.effectiveTags&t > 0
}

func mapDbColumnTypeToGoType(column Column) (goType string, isTime bool) {

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
