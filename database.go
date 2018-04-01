package tablestogo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Database interface for the concrete databases
type Database interface {
	DSN(settings *Settings) string
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

// GeneralDatabase represents a base "class" database - for all other concrete databases
// it implements partly the Database interface
type GeneralDatabase struct {
	GetColumnsOfTableStmt *sqlx.Stmt
	*sqlx.DB
	*Settings
}

func (gdb *GeneralDatabase) Connect(dsn string) (err error) {
	gdb.DB, err = sqlx.Connect(gdb.DbType, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.Settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed:\r\n%v",
			gdb.DbType, gdb.User, gdb.DbName, gdb.Host, gdb.Port, usingPswd, err)
	}

	return gdb.Ping()
}

// Close closes the database connection
func (gdb *GeneralDatabase) Close() error {
	return gdb.Close()
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
