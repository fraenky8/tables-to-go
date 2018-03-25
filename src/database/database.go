package database

import (
	"database/sql"
	"fmt"

	"github.com/fraenky8/tables-to-go/src/settings"
	"github.com/jmoiron/sqlx"
)

var (
	// dbTypeToDriverMap maps the database type to the driver names
	dbTypeToDriverMap = map[string]string{
		"pg":    "postgres",
		"mysql": "mysql",
	}
)

// Table has a name and a set (slice) of columns
type Table struct {
	Name    string `db:"table_name"`
	Columns []Column
}

// Column stores information about a column
type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	Name                   string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	DefaultValue           sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	DatetimePrecision      sql.NullInt64  `db:"datetime_precision"`
	ColumnKey              string         `db:"column_key"`      // mysql specific
	Extra                  string         `db:"extra"`           // mysql specific
	ConstraintName         sql.NullString `db:"constraint_name"` // pg specific
	ConstraintType         sql.NullString `db:"constraint_type"` // pg specific
}

// Database interface for the concrete databases
type Database interface {
	DSN(settings *settings.Settings) string
	Connect() (err error)
	Close() (err error)

	GetTables() (tables []*Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *Table) (err error)

	IsPrimaryKey(column Column) bool
	IsAutoIncrement(column Column) bool
	IsNullable(column Column) bool

	GetStringDatatypes() []string
	IsString(column Column) bool

	GetTextDatatypes() []string
	IsText(column Column) bool

	GetIntegerDatatypes() []string
	IsInteger(column Column) bool

	GetFloatDatatypes() []string
	IsFloat(column Column) bool

	GetTemporalDatatypes() []string
	IsTemporal(column Column) bool

	// TODO pg: bitstrings, enum, range, other special types
	// TODO mysql: bit, enums, set
}

// NewDatabase constructs based on the type in the settings a new Database
func NewDatabase(settings *settings.Settings) Database {

	generalDatabase := &generalDatabase{
		driver:   dbTypeToDriverMap[settings.DbType],
		settings: settings,
	}

	switch settings.DbType {
	case "mysql":
		return &mysql{generalDatabase}
	default: // pg
		return &postgresql{generalDatabase}
	}
}

// generalDatabase represents a base "class" database - for all other concrete databases
// it implements partly the Database interface
type generalDatabase struct {
	driver                string
	db                    *sqlx.DB
	getColumnsOfTableStmt *sqlx.Stmt
	settings              *settings.Settings
}

func (gdb *generalDatabase) connect(dsn string) (err error) {
	gdb.db, err = sqlx.Connect(gdb.driver, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed:\r\n%v",
			gdb.settings.DbType, gdb.settings.User, gdb.settings.DbName, gdb.settings.Host, gdb.settings.Port, usingPswd, err)
	}

	return gdb.db.Ping()
}

// Close closes the database connection
func (gdb *generalDatabase) Close() error {
	return gdb.db.Close()
}

// IsNullable returns true if column is a nullable one
func (gdb *generalDatabase) IsNullable(column Column) bool {
	return column.IsNullable == "YES"
}

// IsStringInSlice checks if needle (string) is in haystack ([]string)
func (gdb *generalDatabase) IsStringInSlice(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
