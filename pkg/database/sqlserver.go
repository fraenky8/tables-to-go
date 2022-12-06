package database

import (
	"fmt"
	"strings"

	"github.com/fraenky8/tables-to-go/pkg/settings"

	// sqlserver database driver
	_ "github.com/denisenkom/go-mssqldb"
)

// Sqlserver implements the Database interface with help of GeneralDatabase.
type Sqlserver struct {
	*GeneralDatabase

	defaultUserName string
}

// NewSqlserver creates a new Sqlserver database.
func NewSqlserver(s *settings.Settings) *Sqlserver {
	return &Sqlserver{
		GeneralDatabase: &GeneralDatabase{
			Settings: s,
			driver:   dbTypeToDriverMap[s.DbType],
		},
		defaultUserName: "sa",
	}
}

// Connect connects to the database by the given data source name (dsn) of the
// concrete database.
func (ss *Sqlserver) Connect() error {
	return ss.GeneralDatabase.Connect(ss.DSN())
}

// DSN creates the DSN String to connect to this database.
func (ss *Sqlserver) DSN() string {
	user := ss.defaultUserName
	if ss.Settings.User != "" {
		user = ss.Settings.User
	}
	return fmt.Sprintf("server=%s;port=%s;user=%s;database=%s;password=%s",
		ss.Settings.Host, ss.Settings.Port, user, ss.Settings.DbName, ss.Settings.Pswd)
}

// GetTables gets all tables for a given schema by name.
func (ss *Sqlserver) GetTables() (tables []*Table, err error) {

	err = ss.Select(&tables, `
		SELECT
            table_name
        FROM
            INFORMATION_SCHEMA.TABLES
        WHERE
            TABLE_SCHEMA = @p1
        ORDER BY TABLE_NAME
	`, ss.Schema)

	if ss.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", ss.Schema)
		}
	}

	return tables, err
}

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the
// columns of a specific table for a given database.
func (ss *Sqlserver) PrepareGetColumnsOfTableStmt() (err error) {

	ss.GetColumnsOfTableStmt, err = ss.Preparex(`
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
		WHERE ic.table_name = @p1
		AND ic.table_schema = @p2
		ORDER BY ic.ordinal_position
	`)

	return err
}

// GetColumnsOfTable executes the statement for retrieving the columns of a
// specific table in a given schema.
func (ss *Sqlserver) GetColumnsOfTable(table *Table) (err error) {

	err = ss.GetColumnsOfTableStmt.Select(&table.Columns, table.Name, ss.Schema)

	if ss.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> schema: %q\r\n", ss.Schema)
		}
	}

	return err
}

// IsPrimaryKey checks if the column belongs to the primary key.
func (ss *Sqlserver) IsPrimaryKey(column Column) bool {
	return strings.Contains(column.ConstraintType.String, "PRIMARY KEY")
}

// IsAutoIncrement checks if the column is an auto_increment column.
func (ss *Sqlserver) IsAutoIncrement(column Column) bool {
	return strings.Contains(column.DefaultValue.String, "IDENTITY")
}

// GetStringDatatypes returns the string datatypes for the Sqlserver database.
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16#character-strings
func (ss *Sqlserver) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"nchar",
		"nvarchar",
	}
}

// IsString returns true if colum is of type string for the Sqlserver database.
func (ss *Sqlserver) IsString(column Column) bool {
	return isStringInSlice(column.DataType, ss.GetStringDatatypes())
}

// GetTextDatatypes returns the text datatypes for the Sqlserver database.
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16#character-strings
func (ss *Sqlserver) GetTextDatatypes() []string {
	return []string{
		"text",
		"ntext",
		"binary",
		"varbinary",
		"image",
	}
}

// IsText returns true if colum is of type text for the Sqlserver database.
func (ss *Sqlserver) IsText(column Column) bool {
	return isStringInSlice(column.DataType, ss.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer datatypes for the Sqlserver database.
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16#exact-numerics
func (ss *Sqlserver) GetIntegerDatatypes() []string {
	return []string{
		"bigint",
		"bit",
		"smallint",
		"int",
		"tinyint",
	}
}

// IsInteger returns true if colum is of type integer for the Sqlserver database.
func (ss *Sqlserver) IsInteger(column Column) bool {
	return isStringInSlice(column.DataType, ss.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float datatypes for the Sqlserver database.
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16#exact-numerics
func (ss *Sqlserver) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
	}
}

// IsFloat returns true if colum is of type float for the Sqlserver database.
func (ss *Sqlserver) IsFloat(column Column) bool {
	return isStringInSlice(column.DataType, ss.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal datatypes for the Sqlserver database.
// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16#date-and-time
func (ss *Sqlserver) GetTemporalDatatypes() []string {
	return []string{
		"date",
		"datetimeoffset",
		"datetime2",
		"smalldatetime",
		"datetime",
		"time",
	}
}

// IsTemporal returns true if colum is of type temporal for the Sqlserver database.
func (ss *Sqlserver) IsTemporal(column Column) bool {
	return isStringInSlice(column.DataType, ss.GetTemporalDatatypes())
}
