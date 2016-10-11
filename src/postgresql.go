package tablestogo

import (
	"fmt"
	"strings"
)

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
			ic.datetime_precision,
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

func (pg *PostgreDatabase) GetStringDatatypes() []string {
	return []string{
		"character varying",
		"varchar",
		"character",
		"char",
	}
}

func (pg *PostgreDatabase) IsString(column Column) bool {
	return IsStringInSlice(column.DataType, pg.GetStringDatatypes())
}

func (pg *PostgreDatabase) GetTextDatatypes() []string {
	return []string{
		"text",
	}
}

func (pg *PostgreDatabase) IsText(column Column) bool {
	return IsStringInSlice(column.DataType, pg.GetTextDatatypes())
}

func (pg *PostgreDatabase) GetIntegerDatatypes() []string {
	return []string{
		"smallint",
		"integer",
		"bigint",
		"smallserial",
		"serial",
		"bigserial",
	}
}

func (pg *PostgreDatabase) IsInteger(column Column) bool {
	return IsStringInSlice(column.DataType, pg.GetIntegerDatatypes())
}

func (pg *PostgreDatabase) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"real",
		"double precision",
	}
}

func (pg *PostgreDatabase) IsFloat(column Column) bool {
	return IsStringInSlice(column.DataType, pg.GetFloatDatatypes())
}

func (pg *PostgreDatabase) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"time with time zone",
		"timestamp with time zone",
		"time without time zone",
		"timestamp without time zone",
		"date",
	}
}

func (pg *PostgreDatabase) IsTemporal(column Column) bool {
	return IsStringInSlice(column.DataType, pg.GetTemporalDatatypes())
}
