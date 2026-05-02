package database

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

var (
	// dbTypeToDriverMap maps the database type to the driver names.
	dbTypeToDriverMap = map[settings.DBType]string{
		settings.DBTypePostgresql: "postgres",
		settings.DBTypeMySQL:      "mysql",
		settings.DBTypeSQLite:     "sqlite",
	}
)

// Database interface for the concrete databases.
type Database interface {
	SQLDriver() *sqlx.DB
	DSN() (string, error)
	Connect(ctx context.Context) error
	Close() error
	Version(ctx context.Context) (string, error)

	GetTables(ctx context.Context, tables ...string) ([]*Table, error)
	PrepareGetColumnsOfTableStmt(ctx context.Context) error
	GetColumnsOfTable(ctx context.Context, table *Table) error

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

// Table has a name and a set (slice) of columns.
type Table struct {
	Name    string `db:"table_name"`
	Comment string `db:"table_comment"`
	Columns []Column
}

// Column stores information about a column.
type Column struct {
	Name                   string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	IsNullable             string         `db:"is_nullable"`
	ColumnKey              string         `db:"column_key"` // mysql specific
	Extra                  string         `db:"extra"`      // mysql specific
	Comment                string         `db:"column_comment"`
	DefaultValue           sql.NullString `db:"column_default"`
	ConstraintName         sql.NullString `db:"constraint_name"` // pg specific
	ConstraintType         sql.NullString `db:"constraint_type"` // pg specific
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	OrdinalPosition        int            `db:"ordinal_position"`
}

// GeneralDatabase represents a base "class" database - for all other concrete
// databases it implements partly the Database interface.
type GeneralDatabase struct {
	GetColumnsOfTableStmt *sqlx.Stmt
	*settings.Settings
	*sqlx.DB

	driver string
}

// New creates a new Database based on the given type in the settings.
func New(s *settings.Settings) Database {

	var db Database

	switch s.DbType {
	case settings.DBTypeSQLite:
		db = NewSQLite(s)
	case settings.DBTypeMySQL:
		db = NewMySQL(s)
	case settings.DBTypePostgresql:
		fallthrough
	default:
		db = NewPostgresql(s)
	}

	return db
}

// Connect establishes a connection to the database with the given DSN.
func (gdb *GeneralDatabase) Connect(ctx context.Context, dsn string) (err error) {
	gdb.DB, err = sqlx.ConnectContext(ctx, gdb.driver, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.Settings.Pswd != "" {
			usingPswd = "yes"
		}
		// Yes, sqlx can return a non-nil and connected DB (but failed ping'ed)
		// in case of a ping error. Hence, lets try to close it.
		_ = gdb.Close()

		return fmt.Errorf(
			"could not connect to database (type=%q, user=%q, database=%q, host='%v:%v', using password: %v): %w",
			gdb.DbType, gdb.User, gdb.DbName, gdb.Host, gdb.Port, usingPswd, err,
		)
	}

	return nil
}

// SQLDriver returns the underlying SQL driver
func (gdb *GeneralDatabase) SQLDriver() *sqlx.DB {
	return gdb.DB
}

// Close closes the database connection.
func (gdb *GeneralDatabase) Close() error {
	if gdb.DB != nil {
		err := gdb.DB.Close()
		gdb.DB = nil
		return err
	}
	return nil
}

// IsNullable returns true if the column is a nullable column.
func (gdb *GeneralDatabase) IsNullable(column Column) bool {
	return column.IsNullable == "YES"
}

// isStringInSlice checks if needle (string) is in haystack ([]string).
func isStringInSlice(needle string, haystack []string) bool {
	return slices.Contains(haystack, needle)
}

func (*GeneralDatabase) andInClause(field string, params []string, args *[]any) string {
	if field == "" || len(params) == 0 {
		return ""
	}

	*args = slices.Grow(*args, len(params))
	for i := range params {
		*args = append(*args, params[i])
	}

	return "AND " + field + " IN (?" + strings.Repeat(",?", len(params)-1) + ")"
}
