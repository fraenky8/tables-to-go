package cli

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

const (
	anyCtx = mock.Anything
)

type mockDB struct {
	mock.Mock
	database.Database
}

func newMockDB(db database.Database) *mockDB {
	return &mockDB{Database: db}
}

func (db *mockDB) Connect(ctx context.Context) error {
	args := db.Called(ctx)
	return args.Error(0)
}

func (db *mockDB) Close() error {
	args := db.Called()
	return args.Error(0)
}

func (db *mockDB) GetTables(ctx context.Context, tables ...string) ([]*database.Table, error) {
	var args mock.Arguments
	if len(tables) == 0 {
		args = db.Called(ctx)
	} else {
		callArgs := make([]any, 0, len(tables)+1)
		callArgs = append(callArgs, ctx)
		for i := range tables {
			callArgs = append(callArgs, tables[i])
		}
		args = db.Called(callArgs...)
	}
	err := args.Error(1)
	if err != nil {
		return nil, err
	}
	return args.Get(0).([]*database.Table), nil
}

func (db *mockDB) PrepareGetColumnsOfTableStmt(ctx context.Context) error {
	args := db.Called(ctx)
	return args.Error(0)
}

func (db *mockDB) GetColumnsOfTable(ctx context.Context, table *database.Table) error {
	args := db.Called(ctx, table)
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

func TestApp_camelCaseString(t *testing.T) {
	t.Parallel()

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
			app := New(settings.New(), nil, nil, os.Stderr)
			actual := app.camelCaseString(tt.input)
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

func TestApp_Run_StringTextColumns(t *testing.T) {
	t.Parallel()

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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName string `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullString `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *string `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *string `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestApp_Run_IntegerColumns(t *testing.T) {
	t.Parallel()

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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName int `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullInt64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *int `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullInt64 `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *int `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestApp_Run_FloatColumns(t *testing.T) {
	t.Parallel()

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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName float64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullFloat64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *float64 `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullFloat64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *float64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestApp_Run_TemporalColumns(t *testing.T) {
	t.Parallel()

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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName time.Time `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullTime `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName *time.Time `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullTime `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName1 *time.Time `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestApp_Run_BooleanColumns(t *testing.T) {
	t.Parallel()

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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName bool `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullBool `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *bool `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullBool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *bool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_UnknownColumns(t *testing.T) {
	t.Parallel()

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

						mdb.
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\ntype TestTable struct {\nColumnName string `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullString `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *string `db:\"column_name\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n\t\"database/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table).
							Return(nil)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package dto\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *string `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							Return(nil)

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
							On("GetTables", anyCtx).
							Return([]*database.Table{table1, table2}, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt", anyCtx).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table1).
							Return(nil)
						mdb.
							On("GetColumnsOfTable", anyCtx, table2).
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

						app := New(s, mdb, w, os.Stderr)

						err := app.Run(t.Context())
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
		{"empty", "", false},
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

func TestApp_formatColumnName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		app      *App
		column   string
		expected string
		isErr    assert.ErrorAssertionFunc
	}{
		{
			desc: "camel case startWithNumber",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				s.Verbose = true // for coverage
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "1fish2fish",
			expected: "X1fish2fish",
			isErr:    assert.NoError,
		},
		{
			desc: "original case startWithNumber",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "1fish2fish",
			expected: "X_1fish2fish",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case startWithNumberFollowedBySpace",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "1 fish",
			expected: "X1_fish",
			isErr:    assert.NoError,
		},
		{
			desc: "original case startWithNumberFollowedBySpace",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "1 fish",
			expected: "X_1_fish",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case containsSpaces",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "my column\twith\nmany\u200bspaces",
			expected: "MyColumnWithManySpaces",
			isErr:    assert.NoError,
		},
		{
			desc: "original case containsSpaces",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "my column\twith\nmany\u200bspaces",
			expected: "My_column_with_many_spaces",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case titleCase",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "MyColumn",
			expected: "MyColumn",
			isErr:    assert.NoError,
		},
		{
			desc: "original case titleCase",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "MyColumn",
			expected: "MyColumn",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case snakeCase",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "my_column",
			expected: "MyColumn",
			isErr:    assert.NoError,
		},
		{
			desc: "original case snakeCase",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "my_column",
			expected: "My_column",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case titleSnake",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "My_Column",
			expected: "MyColumn",
			isErr:    assert.NoError,
		},
		{
			desc: "original case titleSnake",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "My_Column",
			expected: "My_Column",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case numbersOnly",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "123",
			expected: "X123",
			isErr:    assert.NoError,
		},
		{
			desc: "original case numbersOnly",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "123",
			expected: "X_123",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case nonEnglish",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "火",
			expected: "火",
			isErr:    assert.NoError,
		},
		{
			desc: "original case nonEnglish",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "火",
			expected: "火",
			isErr:    assert.NoError,
		},
		{
			desc: "camel case nonEnglishUpper",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatCamelCase
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "Λλ",
			expected: "Λλ",
			isErr:    assert.NoError,
		},
		{
			desc: "original case nonEnglishUpper",
			app: func() *App {
				s := settings.New()
				s.OutputFormat = settings.OutputFormatOriginal
				return New(s, nil, nil, os.Stderr)
			}(),
			column:   "Λλ",
			expected: "Λλ",
			isErr:    assert.NoError,
		},
		{
			desc:     "semicolons returns error",
			app:      New(settings.New(), nil, nil, os.Stderr),
			column:   "MyColumn;",
			expected: "",
			isErr:    assert.Error,
		},
		{
			desc:     "brackets returns error",
			app:      New(settings.New(), nil, nil, os.Stderr),
			column:   "MyColumn()",
			expected: "",
			isErr:    assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual, err := tt.app.formatColumnName(tt.column, "MyTable")
			tt.isErr(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
