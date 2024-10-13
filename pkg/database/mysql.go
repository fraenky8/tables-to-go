package database

import (
	"fmt"
	"strings"

	"github.com/fraenky8/tables-to-go/v2/pkg/settings"

	// MySQL database driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQL implements the Database interface with help of GeneralDatabase.
type MySQL struct {
	*GeneralDatabase

	defaultUserName string
}

// NewMySQL creates a new MySQL database.
func NewMySQL(s *settings.Settings) *MySQL {
	return &MySQL{
		GeneralDatabase: &GeneralDatabase{
			Settings: s,
			driver:   dbTypeToDriverMap[s.DbType],
		},
		defaultUserName: "root",
	}
}

// Connect connects to the database by the given data source name (dsn) of the
// concrete database.
func (mysql *MySQL) Connect() error {
	return mysql.GeneralDatabase.Connect(mysql.DSN())
}

// DSN creates the DSN String to connect to this database.
func (mysql *MySQL) DSN() string {
	user := mysql.defaultUserName
	if mysql.Settings.User != "" {
		user = mysql.Settings.User
	}

	if mysql.Settings.Socket != "" {
		return fmt.Sprintf("%s:%s@unix(%s)/%s",
			user, mysql.Settings.Pswd, mysql.Settings.Socket, mysql.Settings.DbName)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, mysql.Settings.Pswd, mysql.Settings.Host, mysql.Settings.Port, mysql.Settings.DbName)
}

// Version reports the actual version of the MySQL database.
func (mysql *MySQL) Version() (string, error) {
	var version string
	err := mysql.Get(&version, `
		SELECT
			CONCAT(
				@@version, ' ', 
				@@version_comment, ', ', 
				@@version_compile_os, ' ', 
				@@version_compile_machine) as version
	`)
	if err != nil {
		return "", err
	}
	return version, nil
}

// GetDriverImportLibrary returns the golang sql driver specific fot the MySQL database
func (mysql *MySQL) GetDriverImportLibrary() string {
	return `"github.com/go-sql-driver/mysql"`
}

// GetTables gets all tables for a given database by name.
func (mysql *MySQL) GetTables() (tables []*Table, err error) {

	err = mysql.Select(&tables, `
		SELECT table_name AS table_name
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

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the
// columns of a specific table for a given database.
func (mysql *MySQL) PrepareGetColumnsOfTableStmt() (err error) {

	mysql.GetColumnsOfTableStmt, err = mysql.Preparex(`
		SELECT
		  ordinal_position AS ordinal_position,
		  column_name AS column_name,
		  data_type AS data_type,
		  column_default AS column_default,
		  is_nullable AS is_nullable,
		  character_maximum_length AS character_maximum_length,
		  numeric_precision AS numeric_precision,
		  column_key AS column_key,
		  extra AS extra
		FROM information_schema.columns
		WHERE table_name = ?
		AND table_schema = ?
		ORDER BY ordinal_position
	`)

	return err
}

// GetColumnsOfTable executes the statement for retrieving the columns of a
// specific table for a given database.
func (mysql *MySQL) GetColumnsOfTable(table *Table) (err error) {

	err = mysql.GetColumnsOfTableStmt.Select(&table.Columns, table.Name, mysql.DbName)

	if mysql.Settings.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> schema: %q\r\n", mysql.Schema)
			fmt.Printf("> dbName: %q\r\n", mysql.DbName)
		}
	}

	return err
}

// IsPrimaryKey checks if the column belongs to the primary key.
func (mysql *MySQL) IsPrimaryKey(column Column) bool {
	return strings.Contains(column.ColumnKey, "PRI")
}

// IsAutoIncrement checks if the column is an auto_increment column.
func (mysql *MySQL) IsAutoIncrement(column Column) bool {
	return strings.Contains(column.Extra, "auto_increment")
}

// GetStringDatatypes returns the string data types for the MySQL database.
func (mysql *MySQL) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

// IsString returns true if the column is of type string for the MySQL database.
func (mysql *MySQL) IsString(column Column) bool {
	return isStringInSlice(column.DataType, mysql.GetStringDatatypes())
}

// GetTextDatatypes returns the text data types for the MySQL database.
func (mysql *MySQL) GetTextDatatypes() []string {
	return []string{
		"tinytext",
		"text",
		"mediumtext",
		"longtext",
		"tinyblob",
		"blob",
		"mediumblob",
		"longblob",
	}
}

// IsText returns true if column is of type text for the MySQL database.
func (mysql *MySQL) IsText(column Column) bool {
	return isStringInSlice(column.DataType, mysql.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer data types for the MySQL database.
func (mysql *MySQL) GetIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
		"mediumint",
		"int",
		"bigint",
	}
}

// IsInteger returns true if column is of type integer for the MySQL database.
func (mysql *MySQL) IsInteger(column Column) bool {
	return isStringInSlice(column.DataType, mysql.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float data types for the MySQL database.
func (mysql *MySQL) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

// IsFloat returns true if column is of type float for the MySQL database.
func (mysql *MySQL) IsFloat(column Column) bool {
	return isStringInSlice(column.DataType, mysql.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal data types for the MySQL database.
func (mysql *MySQL) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}

// IsTemporal returns true if column is of type temporal for the MySQL database.
func (mysql *MySQL) IsTemporal(column Column) bool {
	return isStringInSlice(column.DataType, mysql.GetTemporalDatatypes())
}
