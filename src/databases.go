package tablestogo

import (
	"database/sql"

	"fmt"

	"github.com/jmoiron/sqlx"
)

type Table struct {
	TableName string `db:"table_name"`
	Columns   []Column
}

type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	ColumnName             string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	ColumnDefault          sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	ColumnKey              string         `db:"column_key"` // mysql specific
	Extra                  string         `db:"extra"`      // mysql specific
}

type Database interface {
	GetTables() (tables []*Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *Table) (err error)
}

type GeneralDatabase struct {
	db                    *sqlx.DB
	GetColumnsOfTableStmt *sqlx.Stmt
	*Settings
}

type PostgreDatabase struct {
	*GeneralDatabase
}

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

func (pg *PostgreDatabase) PrepareGetColumnsOfTableStmt() (err error) {

	pg.GetColumnsOfTableStmt, err = db.Preparex(`
		SELECT
		  ordinal_position,
		  column_name,
		  data_type,
		  column_default,
		  is_nullable,
		  character_maximum_length,
		  numeric_precision
		FROM information_schema.columns
		WHERE table_name = $1
		AND table_schema = $2
		ORDER BY ordinal_position
	`)

	return err
}

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

type MySQLDatabase struct {
	*GeneralDatabase
}

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
