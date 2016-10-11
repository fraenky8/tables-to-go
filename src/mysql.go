package tablestogo

import (
	"fmt"
	"strings"
)

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
		  datetime_precision,
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

func (mysql *MySQLDatabase) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

func (mysql *MySQLDatabase) IsString(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetStringDatatypes())
}

func (mysql *MySQLDatabase) GetTextDatatypes() []string {
	return []string{
		"text",
		"blob",
	}
}

func (mysql *MySQLDatabase) IsText(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetTextDatatypes())
}

func (mysql *MySQLDatabase) GetIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
		"mediumint",
		"int",
		"bigint",
	}
}

func (mysql *MySQLDatabase) IsInteger(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetIntegerDatatypes())
}

func (mysql *MySQLDatabase) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

func (mysql *MySQLDatabase) IsFloat(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetFloatDatatypes())
}

func (mysql *MySQLDatabase) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}

func (mysql *MySQLDatabase) IsTemporal(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetTemporalDatatypes())
}
