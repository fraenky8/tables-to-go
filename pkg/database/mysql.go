package database

import (
	"fmt"
	"strings"

	"github.com/fraenky8/tables-to-go/pkg"
	// Mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

// Mysql implemenmts the Database interface with help of generalDatabase
type Mysql struct {
	*pkg.GeneralDatabase
}

// Connect connects to the database by the given data source name (dsn) of the concrete database
func (mysql *Mysql) Connect() error {
	return mysql.GeneralDatabase.Connect(mysql.DSN(mysql.Settings))
}

// DSN creates the DSN String to connect to this database
func (mysql *Mysql) DSN(settings *pkg.Settings) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		settings.User, settings.Pswd, settings.Host, settings.Port, settings.DbName)
}

// GetTables gets all tables for a given database by name
func (mysql *Mysql) GetTables() (tables []*pkg.Table, err error) {

	err = mysql.Select(&tables, `
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

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the columns of a specific table for a given database
func (mysql *Mysql) PrepareGetColumnsOfTableStmt() (err error) {

	mysql.GetColumnsOfTableStmt, err = mysql.Preparex(`
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

// GetColumnsOfTable executes the statement for retrieving the columns of a specific table for a given database
func (mysql *Mysql) GetColumnsOfTable(table *pkg.Table) (err error) {

	mysql.GetColumnsOfTableStmt.Select(&table.Columns, table.Name, mysql.DbName)

	if mysql.Settings.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> schema: %q\r\n", mysql.Schema)
			fmt.Printf("> dbName: %q\r\n", mysql.DbName)
		}
	}

	return err
}

// IsPrimaryKey checks if column belongs to primary key
func (mysql *Mysql) IsPrimaryKey(column pkg.Column) bool {
	return strings.Contains(column.ColumnKey, "PRI")
}

// IsAutoIncrement checks if column is a auto_increment column
func (mysql *Mysql) IsAutoIncrement(column pkg.Column) bool {
	return strings.Contains(column.Extra, "auto_increment")
}

// GetStringDatatypes returns the string datatypes for the Mysql database
func (mysql *Mysql) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

// IsString returns true if colum is of type string for the Mysql database
func (mysql *Mysql) IsString(column pkg.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetStringDatatypes())
}

// GetTextDatatypes returns the text datatypes for the Mysql database
func (mysql *Mysql) GetTextDatatypes() []string {
	return []string{
		"text",
		"blob",
	}
}

// IsText returns true if colum is of type text for the Mysql database
func (mysql *Mysql) IsText(column pkg.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer datatypes for the Mysql database
func (mysql *Mysql) GetIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
		"mediumint",
		"int",
		"bigint",
	}
}

// IsInteger returns true if colum is of type integer for the Mysql database
func (mysql *Mysql) IsInteger(column pkg.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float datatypes for the Mysql database
func (mysql *Mysql) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

// IsFloat returns true if colum is of type float for the Mysql database
func (mysql *Mysql) IsFloat(column pkg.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal datatypes for the Mysql database
func (mysql *Mysql) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}

// IsTemporal returns true if colum is of type temporal for the Mysql database
func (mysql *Mysql) IsTemporal(column pkg.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetTemporalDatatypes())
}
