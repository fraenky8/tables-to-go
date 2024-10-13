// +build !integration

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

type mockDB struct {
	mock.Mock
	database.Database
}

func newMockDB(db database.Database) *mockDB {
	return &mockDB{Database: db}
}

func (db *mockDB) Connect() error {
	args := db.Called()
	return args.Error(0)
}

func (db *mockDB) Close() error {
	args := db.Called()
	return args.Error(0)
}

func (db *mockDB) GetTables() ([]*database.Table, error) {
	args := db.Called()
	err := args.Error(1)
	if err != nil {
		return nil, err
	}
	return args.Get(0).([]*database.Table), nil
}

func (db *mockDB) PrepareGetColumnsOfTableStmt() error {
	args := db.Called()
	return args.Error(0)
}

func (db *mockDB) GetColumnsOfTable(table *database.Table) error {
	args := db.Called(table)
	return args.Error(0)
}

type mockWriter struct {
	mock.Mock
}

func newMockWriter() *mockWriter {
	return &mockWriter{}
}

func (w *mockWriter) Write(tableName, content string) error {
	args := w.Called(tableName, content)
	return args.Error(0)
}

func TestCamelCaseString(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "empty string returns empty string",
			input:    "",
			expected: "",
		},
		{
			desc:     "single string returns titleized single string",
			input:    "string",
			expected: "String",
		},
		{
			desc:     "multi separated string returns CamelCase string",
			input:    "string_with_separate_sections",
			expected: "StringWithSeparateSections",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := camelCaseString(tt.input)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.input)
		})
	}
}

func TestToInitialisms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "id should be upper case",
			input:    "Id",
			expected: "ID",
		},
		{
			desc:     "id at the end of string should be upper case",
			input:    "userId",
			expected: "userID",
		},
		{
			desc:     "id at the beginning of string should be upper case",
			input:    "Iduser",
			expected: "IDuser",
		},
		{
			desc:     "id in the middle of string should be upper case",
			input:    "userIdprim",
			expected: "userIDprim",
		},
		{
			desc:     "multiple occurrences should be upper case",
			input:    "userIdasJsonWithUrl",
			expected: "userIDasJSONWithURL",
		},
		{
			desc:     "multiple id in the string should be upper case",
			input:    "IduserId",
			expected: "IDuserID",
		},
		{
			desc:     "non replacement in the string should be return original string",
			input:    "name",
			expected: "name",
		},
		{
			desc:     "replacements only in the string should be return original string",
			input:    "IdjsonuRlHtTp",
			expected: "IDJSONURLHTTP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := toInitialisms(tt.input)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.input)
		})
	}
}

func TestRun_StringTextColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := db.GetStringDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName string `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullString `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *string `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *string `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 string `db:\"column_name_1\"`\nColumnName2 sql.NullString `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_IntegerColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := db.GetIntegerDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName int `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullInt64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *int `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullInt64 `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *int `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullInt64 `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 int `db:\"column_name_1\"`\nColumnName2 sql.NullInt64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_FloatColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := db.GetFloatDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName float64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullFloat64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *float64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullFloat64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *float64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullFloat64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 float64 `db:\"column_name_1\"`\nColumnName2 sql.NullFloat64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_TemporalColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := db.GetTemporalDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName time.Time `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullTime `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName *time.Time `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullTime `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName1 *time.Time `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n\t\"time\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullTime `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n\t\"time\"\n)\n\ntype TestTable2 struct {\nColumnName1 time.Time `db:\"column_name_1\"`\nColumnName2 sql.NullTime `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_BooleanColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := []string{"boolean"}

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName bool `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullBool `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *bool `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullBool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *bool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb := newMockDB(db)
						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullBool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 bool `db:\"column_name_1\"`\nColumnName2 sql.NullBool `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_UnknownColumns(t *testing.T) {
	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := settings.New()
			s.DbType = dbType
			db := database.New(s)

			columnTypes := []string{
				"enum",         // MySQL
				"USER-DEFINED", // Postgres
			}

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						mdb := newMockDB(db)

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						//
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName string `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						mdb := newMockDB(db)

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						//
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullString `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						mdb := newMockDB(db)

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						//
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *string `db:\"column_name\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						mdb := newMockDB(db)

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						//
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType
						s.Null = settings.NullTypeNative

						mdb := newMockDB(db)

						table := &database.Table{
							Name: "test_table",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						//
						mdb.
							On("GetTables").
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *string `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := settings.New()
						s.DbType = dbType

						mdb := newMockDB(db)

						table1 := &database.Table{
							Name: "test_table_1",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &database.Table{
							Name: "test_table_2",
							Columns: []database.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}

						mdb.
							On("GetTables").
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table2).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)
						w.
							On(
								"Write",
								"TestTable2",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 string `db:\"column_name_1\"`\nColumnName2 sql.NullString `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						err := Run(s, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestValidVariableName(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    string
		expected bool
	}
	tests := []testCase{
		{"basic", "MyVariable_2", true},
		{"specialChars", "MyVar;iable", false},
		{"brackets", "MyVariabl(e)", false},
		{"nonEnglish", "MyVαriαble", true},
		{"spaces", "My Variable", false},
		{"whitespace", "My		Variable", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if validVariableName(tc.input) != tc.expected {
				t.Errorf("TestValidVariableName(%q) should be %t", tc.input, tc.expected)
			}
		})
	}
}

func TestReplaceSpace(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    rune
		expected rune
	}
	tests := []testCase{
		{"letter", 'a', 'a'},
		{"number", '7', '7'},
		{"nonEnglish", '水', '水'},
		{"space", ' ', '_'},
		{"underscore", '_', '_'},
		{"tab", '\t', '_'},
		{"newline", '\n', '_'},
		{"zeroWidthSpace", '​', '_'},
		{"nonBreakingSpace", ' ', '_'},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := replaceSpace(tc.input)
			if output != tc.expected {
				t.Errorf("replaceSpace(%q) = %q, expected %q", tc.input, output, tc.expected)
			}
		})
	}
}

func TestFormatColumnName(t *testing.T) {
	t.Parallel()

	t.Run("pass", func(t *testing.T) {
		type testCase struct {
			name     string
			input    string
			original string
			camel    string
		}
		tests := []testCase{
			{"startWithNumber", "1fish2fish", "X_1fish2fish", "X1fish2fish"},
			{"containsSpaces", "my column\twith\nmany\u200bspaces", "My_column_with_many_spaces", "MyColumnWithManySpaces"},
			{"titleCase", "MyColumn", "MyColumn", "MyColumn"},
			{"snakeCase", "my_column", "My_column", "MyColumn"},
			{"titleSnake", "My_Column", "My_Column", "MyColumn"},
			{"numbersOnly", "123", "X_123", "X123"},
			{"nonEnglish", "火", "火", "火"},
			{"nonEnglishUpper", "Λλ", "Λλ", "Λλ"},
		}

		camelSettings := settings.New()
		camelSettings.OutputFormat = settings.OutputFormatCamelCase
		t.Run("camelcase", func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					output, err := formatColumnName(camelSettings, tc.input, "MyTable")
					if err != nil {
						t.Error(err)
					} else if output != tc.camel {
						t.Errorf("camelcase format of %q = %q, expected %q", tc.input, output, tc.camel)
					}
				})
			}
		})

		originalSettings := settings.New()
		originalSettings.OutputFormat = settings.OutputFormatOriginal
		t.Run("original", func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					output, err := formatColumnName(originalSettings, tc.input, "MyTable")
					if err != nil {
						t.Error(err)
					} else if output != tc.original {
						t.Errorf("originalCase format of %q = %q, expected %q", tc.input, output, tc.original)
					}
				})
			}
		})
	})

	t.Run("fail", func(t *testing.T) {
		type testCase struct {
			name  string
			input string
		}
		tests := []testCase{
			{"semicolons", "MyColumn;"},
			{"brackets", "MyColumn()"},
		}
		s := settings.New()
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := formatColumnName(s, tc.input, "MyTable")
				if err == nil {
					t.Errorf("formatColumnName(%q) should have thrown error but didn't", tc.input)
				}
			})
		}
	})
}
