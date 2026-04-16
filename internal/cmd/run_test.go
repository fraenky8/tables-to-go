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
		expected *cmdArgs
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
			expected: &cmdArgs{
				Settings: settings.New(),
			},
			isErr: assert.NoError,
		},
		{
			desc: "help flag is parsed",
			args: []string{"tables-to-go", "-help"},
			expected: &cmdArgs{
				Settings: settings.New(),
				Help:     true,
			},
			isErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var stderr bytes.Buffer
			actual, err := newCmdArgs(test.args, &stderr)
			test.isErr(t, err)
			if actual != nil {
				actual.usage = nil
			}
			assert.Equal(t, test.expected, actual)
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
			isErr: func(t assert.TestingT, err error, i ...any) bool {
				return assert.ErrorContains(t, err, "could not connect to database")
			},
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
			cmd := New(version)
			cmd.db = tt.db

			var stdout, stderr bytes.Buffer
			err := cmd.Run(context.Background(), tt.args, &stdout, &stderr)
			tt.isErr(t, err)
			assert.Regexp(t, tt.expectedStdout, stdout.String())
			assert.Regexp(t, tt.expectedStderr, stderr.String())
			tt.db.AssertExpectations(t)
		})
	}
}
