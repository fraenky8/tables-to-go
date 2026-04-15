package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/fraenky8/tables-to-go/v2/pkg/settings"

	// sqlite3 database driver
	_ "modernc.org/sqlite"
)

const (
	defaultBusyTimeout = 5 * time.Second
	defaultCacheSize   = 20 * 1024
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

// Connect connects to the database by the given data source name (dsn) of the
// concrete database.
func (s *SQLite) Connect() (err error) {
	return s.GeneralDatabase.Connect(s.DSN())
}

// DSN creates the DSN String to connect to this database.
// Any Username and Password set in the settings are ignored since SQLite3 does
// not support authentication yet (https://sqlite.org/forum/forumpost/9a4c2a21beb82efd?t=h&unf).
func (s *SQLite) DSN() string {
	normalized := strings.ReplaceAll(s.Settings.DbName, `\`, `/`)

	if !strings.HasPrefix(normalized, "file:") {
		normalized = "file:" + normalized
	}

	u, err := url.Parse(normalized)
	if err != nil {
		return s.Settings.DbName
	}

	q := u.Query()
	if !s.hasDSNParam(q, "busy_timeout") {
		q.Add("_pragma", fmt.Sprintf("busy_timeout(%d)", defaultBusyTimeout.Milliseconds()))
	}
	if !s.hasDSNParam(q, "cache_size") {
		q.Add("_pragma", fmt.Sprintf("cache_size(%d)", defaultCacheSize))
	}
	u.RawQuery = q.Encode()

	return u.RequestURI()
}

func (s *SQLite) hasDSNParam(values url.Values, p string) bool {
	for _, v := range values["_pragma"] {
		if strings.HasPrefix(v, p+`(`) {
			return true
		}
	}
	return false
}

// Version reports the actual version of the Sqlite database.
func (s *SQLite) Version() (string, error) {
	var version string
	err := s.Get(&version, `SELECT sqlite_version()`)
	if err != nil {
		return "", err
	}
	return version, nil
}

// GetTables gets all tables for a given database by name.
func (s *SQLite) GetTables(tables ...string) ([]*Table, error) {

	var args []any
	in := s.andInClause("name", tables, &args)

	var dbTables []*Table
	err := s.Select(&dbTables, `
		SELECT name AS table_name
		FROM sqlite_master
		WHERE type = 'table'
		AND name NOT LIKE 'sqlite?_%' ESCAPE '?'
		`+in+`
	`, args...)

	if s.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> database: %q\r\n", s.DbName)
		}
	}

	return dbTables, err
}

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the
// columns of a specific table for a given database. Unused in Sqlite.
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

// IsPrimaryKey checks if the column belongs to the primary key.
func (s *SQLite) IsPrimaryKey(column Column) bool {
	return column.ColumnKey == "PK"
}

// IsAutoIncrement checks if the column is an auto_increment column.
func (s *SQLite) IsAutoIncrement(column Column) bool {
	return column.ColumnKey == "PK"
}

// GetStringDatatypes returns the string datatypes for the SQLite database.
func (s *SQLite) GetStringDatatypes() []string {
	return []string{
		"text",
	}
}

// IsString returns true if the column is of type string for the SQLite database.
func (s *SQLite) IsString(column Column) bool {
	return isStringInSlice(column.DataType, s.GetStringDatatypes())
}

// GetTextDatatypes returns the text datatypes for the SQLite database.
func (s *SQLite) GetTextDatatypes() []string {
	return []string{
		"text",
	}
}

// IsText returns true if column is of type text for the SQLite database.
func (s *SQLite) IsText(column Column) bool {
	return isStringInSlice(column.DataType, s.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer datatypes for the SQLite database.
func (s *SQLite) GetIntegerDatatypes() []string {
	return []string{
		"integer",
	}
}

// IsInteger returns true if column is of type integer for the SQLite database.
func (s *SQLite) IsInteger(column Column) bool {
	return isStringInSlice(column.DataType, s.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float datatypes for the SQLite database.
func (s *SQLite) GetFloatDatatypes() []string {
	return []string{
		"real",
		"numeric",
	}
}

// IsFloat returns true if column is of type float for the SQLite database.
func (s *SQLite) IsFloat(column Column) bool {
	return isStringInSlice(column.DataType, s.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal datatypes for the SQLite database.
func (s *SQLite) GetTemporalDatatypes() []string {
	return []string{}
}

// IsTemporal returns true if column is of type temporal for the SQLite database.
func (s *SQLite) IsTemporal(_ Column) bool {
	return false
}

// GetTemporalDriverDataType returns the time data type specific for the Sqlite database.
func (s *SQLite) GetTemporalDriverDataType() string {
	return ""
}
