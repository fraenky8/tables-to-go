package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fraenky8/tables-to-go/pkg/config"
	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
)

type mockMySQLDb struct {
	mock.Mock

	*database.MySQL

	tables []*database.Table
}

func newMockMySQLDb(settings *config.Settings) *mockMySQLDb {
	return &mockMySQLDb{
		MySQL: database.NewMySQL(settings),
	}
}

func (db *mockMySQLDb) Connect() (err error) {
	db.Called()
	return nil
}

func (db *mockMySQLDb) Close() (err error) {
	db.Called()
	return nil
}

func (db *mockMySQLDb) GetTables() (tables []*database.Table, err error) {
	db.Called()
	return db.tables, nil
}

func (db *mockMySQLDb) PrepareGetColumnsOfTableStmt() (err error) {
	db.Called()
	return nil
}

func (db *mockMySQLDb) GetColumnsOfTable(table *database.Table) (err error) {
	db.Called(table)
	return nil
}

type mockWriter struct {
	mock.Mock
}

func newMockWriter() *mockWriter {
	return &mockWriter{}
}

func (w *mockWriter) Write(tableName string, content string) error {
	w.Called(tableName, content)
	return nil
}

func Test_toInitialisms(t *testing.T) {
	tests := []struct {
		desc     string
		intput   string
		expected string
	}{
		{
			desc:     "id should be upper case",
			intput:   "Id",
			expected: "ID",
		},
		{
			desc:     "id at the end of string should be upper case",
			intput:   "userId",
			expected: "userID",
		},
		{
			desc:     "id at the beginning of string should be upper case",
			intput:   "Iduser",
			expected: "IDuser",
		},
		{
			desc:     "id in the middle of string should be upper case",
			intput:   "userIdprim",
			expected: "userIDprim",
		},
		{
			desc:     "multiple occurrences should be upper case",
			intput:   "userIdasJsonWithUrl",
			expected: "userIDasJSONWithURL",
		},
		{
			desc:     "multiple id in the string should be upper case",
			intput:   "IduserId",
			expected: "IDuserID",
		},
		{
			desc:     "non replacement in the string should be return original string",
			intput:   "name",
			expected: "name",
		},
		{
			desc:     "replacements only in the string should be return original string",
			intput:   "IdjsonuRlHtTp",
			expected: "IDJSONURLHTTP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := toInitialisms(tt.intput)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.intput)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *config.Settings
		db       func(*config.Settings) database.Database
		writer   func() output.Writer
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc: "mysql: single table with integer column",
			settings: func() *config.Settings {
				s := config.NewSettings()
				s.DbType = config.DbTypeMySQL
				return s
			},
			db: func(settings *config.Settings) database.Database {
				db := newMockMySQLDb(settings)

				table := &database.Table{
					Name: "test_table",
					Columns: []database.Column{
						{
							OrdinalPosition: 1,
							Name:            "user_id",
							DataType:        "integer",
						},
					},
				}
				db.tables = append(db.tables, table)

				db.
					On("GetTables").
					Return(db.tables, nil)
				db.
					On("PrepareGetColumnsOfTableStmt").
					Return(nil)
				db.
					On("GetColumnsOfTable", table)

				return db
			},
			writer: func() output.Writer {
				w := newMockWriter()
				w.
					On(
						"Write",
						"TestTable",
						"package dto\n\ntype TestTable struct {\nUserID sql.NullString `db:\"user_id\"`\n}",
					)
				return w
			},
			isError: assert.NoError,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings()
			db := test.db(s)
			w := test.writer()

			err := Run(s, db, w)
			test.isError(t, err)
		})
	}
}
