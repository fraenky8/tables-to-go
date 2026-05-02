package cmd

import (
	"bytes"
	"context"
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

func TestNewCmdArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		args     []string
		expected *Args
		isErr    assert.ErrorAssertionFunc
	}{
		{
			desc:     "unknown flag returns parse sentinel error",
			args:     []string{"tables-to-go", "-unknown"},
			expected: nil,
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorIs(t, err, ErrFlagParse)
			},
		},
		{
			desc: "empty args defaults to binary name",
			args: []string{},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "help flag is parsed",
			args: []string{"tables-to-go", "-help"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.ResolveTags()
					return s
				}(),
				Help: true,
			},
			isErr: assert.NoError,
		},
		{
			desc: "single tags value is parsed",
			args: []string{"tables-to-go", "-tag", "structable"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.Tags = settings.StringsFlag{"structable"}
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "gorm model flag is parsed",
			args: []string{"tables-to-go", "-gorm-model"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.IsGormModel = true
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "gen header flag is parsed",
			args: []string{"tables-to-go", "-gen-header"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.GenHeader = true
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "comments flag without value is parsed as line mode",
			args: []string{"tables-to-go", "-comments"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.Comments = settings.CommentsModeLine
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "comments flag with inline value is parsed",
			args: []string{"tables-to-go", "-comments=inline"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.Comments = settings.CommentsModeInline
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "comments flag with empty value is parsed as line mode",
			args: []string{"tables-to-go", "-comments="},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.Comments = settings.CommentsModeLine
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "comments flag with explicit off value is parsed",
			args: []string{"tables-to-go", "-comments=off"},
			expected: &Args{
				Settings: func() *settings.Settings {
					s := settings.New()
					s.Comments = settings.CommentsModeOff
					s.ResolveTags()
					return s
				}(),
			},
			isErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var stderr bytes.Buffer
			actual, err := NewArgs(test.args, &stderr)
			test.isErr(t, err)
			if actual != nil {
				actual.usage = nil
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNewCmdArgs_printLegacyTagsWarning(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		args     []string
		expected assert.ComparisonAssertionFunc
	}{
		{
			desc:     "no legacy tags flag emits no warning",
			args:     []string{"tables-to-go", "-tag", "db"},
			expected: assert.NotContains,
		},
		{
			desc:     "deprecated tags flag emits warning",
			args:     []string{"tables-to-go", "-tags-structable"},
			expected: assert.Contains,
		},
		{
			desc:     "multiple deprecated tags flags emit warning",
			args:     []string{"tables-to-go", "-tags-no-db", "-tags-structable-only"},
			expected: assert.Contains,
		},
		{
			desc:     "tags-no-db alone emits no warning",
			args:     []string{"tables-to-go", "-tags-no-db"},
			expected: assert.NotContains,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var stderr bytes.Buffer

			_, err := NewArgs(test.args, &stderr)
			assert.NoError(t, err)
			test.expected(t, stderr.String(), legacyTagsDeprecationWarning)
		})
	}
}

func TestNewCmdArgs_printRecorderWithoutStructableWarning(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		args     []string
		expected assert.ComparisonAssertionFunc
	}{
		{
			desc:     "recorder with default db tags emits warning",
			args:     []string{"tables-to-go", "-structable-recorder"},
			expected: assert.Contains,
		},
		{
			desc:     "recorder with explicit structable tag emits no warning",
			args:     []string{"tables-to-go", "-structable-recorder", "-tag", "structable"},
			expected: assert.NotContains,
		},
		{
			desc:     "recorder with legacy structable tag emits no warning",
			args:     []string{"tables-to-go", "-structable-recorder", "-tags-structable"},
			expected: assert.NotContains,
		},
		{
			desc:     "no recorder emits no warning",
			args:     []string{"tables-to-go", "-tag", "db"},
			expected: assert.NotContains,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var stderr bytes.Buffer

			_, err := NewArgs(test.args, &stderr)
			assert.NoError(t, err)
			test.expected(t, stderr.String(), recorderWithoutStructableWarning)
		})
	}
}

func TestNewCmdArgs_printGormModelWithoutGormWarning(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		args     []string
		expected assert.ComparisonAssertionFunc
	}{
		{
			desc:     "gorm model with default db tags emits warning",
			args:     []string{"tables-to-go", "-gorm-model"},
			expected: assert.Contains,
		},
		{
			desc:     "gorm model with explicit gorm tag emits no warning",
			args:     []string{"tables-to-go", "-gorm-model", "-tag", "gorm"},
			expected: assert.NotContains,
		},
		{
			desc:     "gorm tag without gorm model emits no warning",
			args:     []string{"tables-to-go", "-tag", "gorm"},
			expected: assert.NotContains,
		},
		{
			desc:     "no gorm model emits no warning",
			args:     []string{"tables-to-go", "-tag", "db"},
			expected: assert.NotContains,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var stderr bytes.Buffer

			_, err := NewArgs(test.args, &stderr)
			assert.NoError(t, err)
			test.expected(t, stderr.String(), gormModelWithoutGormWarning)
		})
	}
}

func Test_Run(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc           string
		args           []string
		db             *mockDB
		expectedStdout string
		expectedStderr string
		isErr          assert.ErrorAssertionFunc
	}{
		{
			desc:           "parse error returns sentinel error",
			args:           []string{"tables-to-go", "-unknown"},
			db:             newMockDB(database.New(settings.New())),
			expectedStdout: "^$",
			expectedStderr: `.*flag provided but not defined: -unknown.*`,
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorIs(t, err, ErrFlagParse)
			},
		},
		{
			desc:           "help returns without error",
			args:           []string{"tables-to-go", "-help"},
			db:             newMockDB(database.New(settings.New())),
			expectedStdout: "^$",
			expectedStderr: `.*Usage of tables-to-go:.*`,
			isErr:          assert.NoError,
		},
		{
			desc:           "version writes output and returns without error",
			args:           []string{"tables-to-go", "-version"},
			db:             newMockDB(database.New(settings.New())),
			expectedStdout: `.*tables-to-go/v0\.0\.0-.*on Jan 02 2006 15:04:05.*`,
			expectedStderr: "^$",
			isErr:          assert.NoError,
		},
		{
			desc:           "verify returns error for invalid output path",
			args:           []string{"tables-to-go", "-of", "_does_not_exist_/output"},
			db:             newMockDB(database.New(settings.New())),
			expectedStdout: "^$",
			expectedStderr: "^$",
			isErr:          assert.Error,
		},
		{
			desc: "connect error does not trigger close",
			args: []string{"tables-to-go"},
			db: func() *mockDB {
				db := newMockDB(database.New(settings.New()))
				db.
					On("Connect", anyCtx).
					Return(assert.AnError)
				return db
			}(),
			expectedStdout: "^$",
			expectedStderr: "^$",
			isErr:          assert.Error,
		},
		{
			desc: "run error prints running header",
			args: []string{"tables-to-go"},
			db: func() *mockDB {
				db := newMockDB(database.New(settings.New()))
				db.
					On("Connect", anyCtx).
					Return(nil)
				db.
					On("GetTables", anyCtx).
					Return([]*database.Table{}, assert.AnError)
				db.
					On("Close").
					Return(nil)
				return db
			}(),
			expectedStdout: "^$",
			expectedStderr: `.*running for.*`,
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorContains(t, err, "run error")
			},
		},
		{
			desc: "run error and close error are combined",
			args: []string{"tables-to-go"},
			db: func() *mockDB {
				db := newMockDB(database.New(settings.New()))
				db.
					On("Connect", anyCtx).
					Return(nil)
				db.
					On("GetTables", anyCtx).
					Return([]*database.Table{}, assert.AnError)
				db.
					On("Close").
					Return(assert.AnError)
				return db
			}(),
			expectedStdout: "^$",
			expectedStderr: `.*running for.*`,
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorContains(t, err, "run error") &&
					assert.ErrorContains(t, err, "could not close database") &&
					assert.ErrorContains(t, err, "previous error")
			},
		},
		{
			desc: "close error after successful run",
			args: []string{"tables-to-go"},
			db: func() *mockDB {
				db := newMockDB(database.New(settings.New()))
				db.
					On("Connect", anyCtx).
					Return(nil)
				db.
					On("GetTables", anyCtx).
					Return([]*database.Table{}, nil)
				db.
					On("PrepareGetColumnsOfTableStmt", anyCtx).
					Return(nil)
				db.
					On("Close").
					Return(assert.AnError)
				return db
			}(),
			expectedStdout: "^$",
			expectedStderr: `.*done!.*`,
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorContains(t, err, "could not close database")
			},
		},
		{
			desc: "successful run writes done output",
			args: []string{"tables-to-go"},
			db: func() *mockDB {
				db := newMockDB(database.New(settings.New()))
				db.
					On("Connect", anyCtx).
					Return(nil)
				db.
					On("GetTables", anyCtx).
					Return([]*database.Table{}, nil)
				db.
					On("PrepareGetColumnsOfTableStmt", anyCtx).
					Return(nil)
				db.
					On("Close").
					Return(nil)
				return db
			}(),
			expectedStdout: "^$",
			expectedStderr: `.*done!.*`,
			isErr:          assert.NoError,
		},
	}

	version := VersionInfo{
		Revision:       "master",
		VersionTag:     "v0.0.0",
		BuildTimestamp: "Jan 02 2006 15:04:05",
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			cmd := New(version, tt.db)

			var stdout, stderr bytes.Buffer
			err := cmd.Run(context.Background(), tt.args, &stdout, &stderr)
			tt.isErr(t, err)
			assert.Regexp(t, tt.expectedStdout, stdout.String())
			assert.Regexp(t, tt.expectedStderr, stderr.String())
			tt.db.AssertExpectations(t)
		})
	}
}
