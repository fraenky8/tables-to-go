package tablestogo

import (
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

// NewDatabase constructs based on the type in the settings a new Database
func NewDatabase(settings *Settings) Database {

	generalDatabase := &GeneralDatabase{
		driver:   dbTypeToDriverMap[settings.DbType],
		settings: settings,
	}

	switch settings.DbType {
	case "mysql":
		return &MySQLDatabase{generalDatabase}
	default: // pg
		return &PostgreDatabase{generalDatabase}
	}
}

// GeneralDatabase represents a base "class" database - for all other concrete databases
// it implements partly the Database interface
type GeneralDatabase struct {
	driver                string
	db                    *sqlx.DB
	getColumnsOfTableStmt *sqlx.Stmt
	settings              *Settings
}

func (gdb *GeneralDatabase) connect(dsn string) (err error) {
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

func (gdb *GeneralDatabase) Close() error {
	return gdb.db.Close()
}

// IsNullable returns true if column is a nullable one
func (gdb *GeneralDatabase) IsNullable(column Column) bool {
	return column.IsNullable == "YES"
}
