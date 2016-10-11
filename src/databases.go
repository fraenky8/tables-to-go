package tablestogo

import (
	"fmt"

	"strings"

	"github.com/jmoiron/sqlx"
)

// database interface for the concrete databases
type Database interface {
	GetTables() (tables []*Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *Table) (err error)
	IsPrimaryKey(column Column) bool
	IsAutoIncrement(column Column) bool
	CreateDataSourceName(settings *Settings) string
}

// a generic database - like a parent/base class of all other concrete databases
type GeneralDatabase struct {
	db                    *sqlx.DB
	GetColumnsOfTableStmt *sqlx.Stmt
	*Settings
}

// concrete database support for PostgreSQL
// PostgreSQL satisfy the database interface
type PostgreDatabase struct {
	*GeneralDatabase
}

// gets all tables for a given schema
func (pg *PostgreDatabase) GetTables() (tables []*Table, err error) {

	err = db.Select(&tables, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		AND table_schema = $1
		ORDER BY table_name
	`, pg.Schema)

	if pg.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", pg.Schema)
		}
	}

	return tables, err
}

// prepares the statement for retrieving the columns of a specific table in a given schema
func (pg *PostgreDatabase) PrepareGetColumnsOfTableStmt() (err error) {

	pg.GetColumnsOfTableStmt, err = db.Preparex(`
		SELECT
			ic.ordinal_position,
			ic.column_name,
			ic.data_type,
			ic.column_default,
			ic.is_nullable,
			ic.character_maximum_length,
			ic.numeric_precision,
			itc.constraint_name,
			itc.constraint_type
		FROM information_schema.columns AS ic
			LEFT JOIN information_schema.key_column_usage AS ikcu ON ic.table_name = ikcu.table_name
			AND ic.table_schema = ikcu.table_schema
			AND ic.column_name = ikcu.column_name
			LEFT JOIN information_schema.table_constraints AS itc ON ic.table_name = itc.table_name
			AND ic.table_schema = itc.table_schema
			AND ikcu.constraint_name = itc.constraint_name
		WHERE ic.table_name = $1
		AND ic.table_schema = $2
		ORDER BY ic.ordinal_position
	`)

	return err
}

// executes the statement for retrieving the columns of a specific table in a given schema
func (pg *PostgreDatabase) GetColumnsOfTable(table *Table) (err error) {

	pg.GetColumnsOfTableStmt.Select(&table.Columns, table.TableName, pg.Schema)

	if pg.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.TableName)
			fmt.Printf("> schema: %q\r\n", pg.Schema)
		}
	}

	return err
}

// checks if column belongs to primary key
func (pg *PostgreDatabase) IsPrimaryKey(column Column) bool {
	return strings.Contains(column.ConstraintType.String, "PRIMARY KEY")
}

// checks if column is a serial column
func (pg *PostgreDatabase) IsAutoIncrement(column Column) bool {
	return strings.Contains(column.ColumnDefault.String, "nextval")
}

// creates the DSN String to connect to this database
func (pg *PostgreDatabase) CreateDataSourceName(settings *Settings) string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		settings.Host, settings.Port, settings.User, settings.DbName, settings.Pswd)
}

// concrete database support for MySQL
// MySQL satisfy the database interface
type MySQLDatabase struct {
	*GeneralDatabase
}

// gets all tables for a given database by name
func (mysql *MySQLDatabase) GetTables() (tables []*Table, err error) {

	err = db.Select(&tables, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		AND table_schema = ?
		ORDER BY table_name
	`, mysql.DbName)

	if mysql.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", mysql.DbName)
		}
	}

	return tables, err
}

// prepares the statement for retrieving the columns of a specific table for a given database
func (mysql *MySQLDatabase) PrepareGetColumnsOfTableStmt() (err error) {

	mysql.GetColumnsOfTableStmt, err = db.Preparex(`
		SELECT
		  ordinal_position,
		  column_name,
		  data_type,
		  column_default,
		  is_nullable,
		  character_maximum_length,
		  numeric_precision,
		  column_key,
		  extra
		FROM information_schema.columns
		WHERE table_name = ?
		AND table_schema = ?
		ORDER BY ordinal_position
	`)

	return err
}

// executes the statement for retrieving the columns of a specific table for a given database
func (mysql *MySQLDatabase) GetColumnsOfTable(table *Table) (err error) {

	mysql.GetColumnsOfTableStmt.Select(&table.Columns, table.TableName, mysql.DbName)

	if mysql.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.TableName)
			fmt.Printf("> schema: %q\r\n", mysql.Schema)
			fmt.Printf("> dbName: %q\r\n", mysql.DbName)
		}
	}

	return err
}

// checks if column belongs to primary key
func (mysql *MySQLDatabase) IsPrimaryKey(column Column) bool {
	return strings.Contains(column.ColumnKey, "PRI")
}

// checks if column is a auto_increment column
func (mysql *MySQLDatabase) IsAutoIncrement(column Column) bool {
	return strings.Contains(column.Extra, "auto_increment")
}

// creates the DSN String to connect to this database
func (mysql *MySQLDatabase) CreateDataSourceName(settings *Settings) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		settings.User, settings.Pswd, settings.Host, settings.Port, settings.DbName)
}
