package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/fraenky8/tables-to-go/pkg/settings"
)

// SQLite implements the Database interface with help of GeneralDatabase.
type SQLite struct {
	*GeneralDatabase
}

// NewSQLite creates a new SQLite database.
func NewSQLite(s *settings.Settings) *SQLite {
	return &SQLite{
		GeneralDatabase: &GeneralDatabase{
			Settings: s,
			driver:   dbTypeToDriverMap[s.DbType],
		},
	}
}

// Connect connects to the database by the given data source name (DSN).
func (s *SQLite) Connect() (err error) {
	return s.GeneralDatabase.Connect(s.DSN())
}

// DSN creates the data source name string to connect to this database.
func (s *SQLite) DSN() string {
	if s.Settings.User == "" && s.Settings.Pswd == "" {
		return fmt.Sprintf("%v", s.Settings.DbName)
	}

	u, err := url.Parse(s.DbName)
	if err != nil {
		return fmt.Sprintf("%v", s.Settings.DbName)
	}

	query := u.Query()
	query.Set("_auth_user", s.Settings.User)
	query.Set("_auth_pass", s.Settings.Pswd)
	u.RawQuery = query.Encode()

	// SQLite driver expects a empty `_auth` request param
	return strings.ReplaceAll(u.RequestURI(), "_auth=&", "_auth&")
}

// Version reports the actual version of the Sqlite database.
func (s *SQLite) Version() (string, error) {
	var version string
	err := s.Get(&version, `SELECT * FROM PRAGMA user_version`)
	if err != nil {
		return "", err
	}
	return version, nil
}

// GetDriverImportLibrary returns the golang sql driver specific for the Sqlite database.
func (s *SQLite) GetDriverImportLibrary() string {
	return `"github.com/mattn/go-sqlite3"`
}

// GetTables gets all tables for a given database by name.
func (s *SQLite) GetTables() (tables []*Table, err error) {

	err = s.Select(&tables, `
		SELECT name AS table_name
		FROM sqlite_master
		WHERE type = 'table'
		AND name NOT LIKE 'sqlite?_%' escape '?'
	`)

	if s.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> database: %q\r\n", s.DbName)
		}
	}

	return tables, err
}

// PrepareGetColumnsOfTableStmt is unused in Sqlite.
func (s *SQLite) PrepareGetColumnsOfTableStmt() (err error) {
	return nil
}

// GetColumnsOfTable executes the statement for retrieving the columns of a
// specific table for a given database.
func (s *SQLite) GetColumnsOfTable(table *Table) (err error) {

	rows, err := s.Queryx(`
		SELECT * 
		FROM PRAGMA_TABLE_INFO('` + table.Name + `')
	`)
	if err != nil {
		if s.Verbose {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> database: %q\r\n", s.DbName)
		}
		return err
	}

	type column struct {
		CID          int            `db:"cid"`
		Name         string         `db:"name"`
		DataType     string         `db:"type"`
		NotNull      int            `db:"notnull"`
		DefaultValue sql.NullString `db:"dflt_value"`
		PrimaryKey   int            `db:"pk"`
	}

	for rows.Next() {
		var col column
		err = rows.StructScan(&col)
		if err != nil {
			return err
		}

		isNullable := "YES"
		if col.NotNull == 1 {
			isNullable = "NO"
		}

		isPrimaryKey := ""
		if col.PrimaryKey == 1 {
			isPrimaryKey = "PK"
		}

		table.Columns = append(table.Columns, Column{
			OrdinalPosition:        col.CID,
			Name:                   col.Name,
			DataType:               col.DataType,
			DefaultValue:           col.DefaultValue,
			IsNullable:             isNullable,
			CharacterMaximumLength: sql.NullInt64{},
			NumericPrecision:       sql.NullInt64{},
			// reuse mysql column_key as primary key indicator
			ColumnKey:      isPrimaryKey,
			Extra:          "",
			ConstraintName: sql.NullString{},
			ConstraintType: sql.NullString{},
		})
	}

	return nil
}

// IsPrimaryKey checks if column belongs to primary key.
func (s *SQLite) IsPrimaryKey(column Column) bool {
	return column.ColumnKey == "PK"
}

// IsAutoIncrement checks if column is an auto_increment column.
func (s *SQLite) IsAutoIncrement(column Column) bool {
	return column.ColumnKey == "PK"
}

// GetStringDatatypes returns the string data types for the Sqlite database.
func (s *SQLite) GetStringDatatypes() []string {
	return []string{
		"text",
	}
}

// IsString returns true if column is of type string for the Sqlite database.
func (s *SQLite) IsString(column Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetStringDatatypes())
}

// GetTextDatatypes returns the text data types for the Sqlite database.
func (s *SQLite) GetTextDatatypes() []string {
	return []string{
		"text",
	}
}

// IsText returns true if column is of type text for the Sqlite database.
func (s *SQLite) IsText(column Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer data types for the Sqlite database.
func (s *SQLite) GetIntegerDatatypes() []string {
	return []string{
		"integer",
	}
}

// IsInteger returns true if column is of type integer for the Sqlite database.
func (s *SQLite) IsInteger(column Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float data types for the Sqlite database.
func (s *SQLite) GetFloatDatatypes() []string {
	return []string{
		"real",
		"numeric",
	}
}

// IsFloat returns true if column is of type float for the Sqlite database.
func (s *SQLite) IsFloat(column Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal data types for the Sqlite database.
func (s *SQLite) GetTemporalDatatypes() []string {
	return []string{}
}

// IsTemporal returns true if column is of type temporal for the Sqlite database.
func (s *SQLite) IsTemporal(column Column) bool {
	return false
}

// GetTemporalDriverDataType returns the time data type specific for the Sqlite database.
func (s *SQLite) GetTemporalDriverDataType() string {
	return ""
}
