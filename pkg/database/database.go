package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/fraenky8/tables-to-go/pkg/settings"
)

var (
	// dbTypeToDriverMap maps the database type to the driver names
	dbTypeToDriverMap = map[settings.DbType]string{
		settings.DbTypePostgresql: "postgres",
		settings.DbTypeMySQL:      "mysql",
		settings.DbTypeSQLite:     "sqlite3",
	}
)

// Database interface for the concrete databases
type Database interface {
	SQLDriver() *sqlx.DB
	DSN() string
	Connect() error
	Close() error
	Version() (string, error)
	GetDriverImportLibrary() string

	GetTables() ([]*Table, error)
	PrepareGetColumnsOfTableStmt() error
	GetColumnsOfTable(table *Table) error

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
	GetTemporalDriverDataType() string

	// TODO pg: bitstrings, enum, range, other special types
	// TODO mysql: bit, enums, set
}

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
	ColumnKey              string         `db:"column_key"`      // mysql specific
	Extra                  string         `db:"extra"`           // mysql specific
	ConstraintName         sql.NullString `db:"constraint_name"` // pg specific
	ConstraintType         sql.NullString `db:"constraint_type"` // pg specific
}

// GeneralDatabase represents a base "class" database - for all other concrete databases
// it implements partly the Database interface
type GeneralDatabase struct {
	GetColumnsOfTableStmt *sqlx.Stmt
	driver                string

	*settings.Settings
	*sqlx.DB
}

// New creates a new Database based on the given type in the settings.
func New(s *settings.Settings) Database {

	var db Database

	switch s.DbType {
	case settings.DbTypeSQLite:
		db = NewSQLite(s)
	case settings.DbTypeMySQL:
		db = NewMySQL(s)
	case settings.DbTypePostgresql:
		fallthrough
	default:
		db = NewPostgresql(s)
	}

	return db
}

// Connect establishes a connection to the database with the given DSN.
// It pings the database to ensure it is reachable.
func (gdb *GeneralDatabase) Connect(dsn string) (err error) {
	gdb.DB, err = sqlx.Connect(gdb.driver, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.Settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf(
			"could not connect to database (type=%q, user=%q, database=%q, host='%v:%v', using password: %v): %v",
			gdb.DbType, gdb.User, gdb.DbName, gdb.Host, gdb.Port, usingPswd, err,
		)
	}

	return gdb.Ping()
}

// SQLDriver returns the underlyig SQL driver
func (gdb *GeneralDatabase) SQLDriver() *sqlx.DB {
	return gdb.DB
}

// Close closes the database connection
func (gdb *GeneralDatabase) Close() error {
	return gdb.DB.Close()
}

// IsNullable returns true if column is a nullable one
func (gdb *GeneralDatabase) IsNullable(column Column) bool {
	return column.IsNullable == "YES"
}

// IsStringInSlice checks if needle (string) is in haystack ([]string)
func (gdb *GeneralDatabase) IsStringInSlice(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
