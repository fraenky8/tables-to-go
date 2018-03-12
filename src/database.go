package tablestogo

import "github.com/jmoiron/sqlx"

// Database interface for the concrete databases
type Database interface {
	GetTables() (tables []*Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *Table) (err error)
	CreateDataSourceName(settings *Settings) string

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

// GeneralDatabase represents a base "class" database - for all other concrete databases
type GeneralDatabase struct {
	Database
	Db                    *sqlx.DB
	GetColumnsOfTableStmt *sqlx.Stmt
	settings              *Settings
}

// IsNullable returns true if column is a nullable one
func (gdb *GeneralDatabase) IsNullable(column Column) bool {
	return column.IsNullable == "YES"
}
