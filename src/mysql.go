package tablestogo

import (
	"fmt"
	"strings"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQLDatabase implemenmts the Database interface with help of GeneralDatabase
type MySQLDatabase struct {
	*GeneralDatabase
}

// Connect connects to the database by the given data source name (dsn) of the concrete database
func (mysql *MySQLDatabase) Connect() error {
	return mysql.connect(mysql.DSN(mysql.settings))
}

// GetTables gets all tables for a given database by name
func (mysql *MySQLDatabase) GetTables() (tables []*Table, err error) {

	err = mysql.db.Select(&tables, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		AND table_schema = ?
		ORDER BY table_name
	`, mysql.settings.DbName)

	if mysql.settings.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", mysql.settings.DbName)
		}
	}

	return tables, err
}

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the columns of a specific table for a given database
func (mysql *MySQLDatabase) PrepareGetColumnsOfTableStmt() (err error) {

	mysql.getColumnsOfTableStmt, err = mysql.db.Preparex(`
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

// GetColumnsOfTable executes the statement for retrieving the columns of a specific table for a given database
func (mysql *MySQLDatabase) GetColumnsOfTable(table *Table) (err error) {

	mysql.getColumnsOfTableStmt.Select(&table.Columns, table.TableName, mysql.settings.DbName)

	if mysql.settings.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.TableName)
			fmt.Printf("> schema: %q\r\n", mysql.settings.Schema)
			fmt.Printf("> dbName: %q\r\n", mysql.settings.DbName)
		}
	}

	return err
}

// IsPrimaryKey checks if column belongs to primary key
func (mysql *MySQLDatabase) IsPrimaryKey(column Column) bool {
	return strings.Contains(column.ColumnKey, "PRI")
}

// IsAutoIncrement checks if column is a auto_increment column
func (mysql *MySQLDatabase) IsAutoIncrement(column Column) bool {
	return strings.Contains(column.Extra, "auto_increment")
}

// DSN creates the DSN String to connect to this database
func (mysql *MySQLDatabase) DSN(settings *Settings) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		settings.User, settings.Pswd, settings.Host, settings.Port, settings.DbName)
}

// GetStringDatatypes returns the string datatypes for the mysql database
func (mysql *MySQLDatabase) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

// IsString returns true if colum is of type string for the mysql database
func (mysql *MySQLDatabase) IsString(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetStringDatatypes())
}

// GetTextDatatypes returns the text datatypes for the mysql database
func (mysql *MySQLDatabase) GetTextDatatypes() []string {
	return []string{
		"text",
		"blob",
	}
}

// IsText returns true if colum is of type text for the mysql database
func (mysql *MySQLDatabase) IsText(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer datatypes for the mysql database
func (mysql *MySQLDatabase) GetIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
		"mediumint",
		"int",
		"bigint",
	}
}

// IsInteger returns true if colum is of type integer for the mysql database
func (mysql *MySQLDatabase) IsInteger(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float datatypes for the mysql database
func (mysql *MySQLDatabase) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

// IsFloat returns true if colum is of type float for the mysql database
func (mysql *MySQLDatabase) IsFloat(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal datatypes for the mysql database
func (mysql *MySQLDatabase) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}

// IsTemporal returns true if colum is of type temporal for the mysql database
func (mysql *MySQLDatabase) IsTemporal(column Column) bool {
	return IsStringInSlice(column.DataType, mysql.GetTemporalDatatypes())
}
