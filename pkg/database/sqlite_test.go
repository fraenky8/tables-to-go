package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

func TestSQLite_DSN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings *settings.Settings
		expected string
		isErr    assert.ErrorAssertionFunc
	}{
		{
			desc: "an invalid URL query param returns error",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=\x00"
				return s
			}(),
			expected: "",
			isErr:    assert.Error,
		},
		{
			desc: "plain db file without query params",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db"
				return s
			}(),
			expected: "path/to/a/file.db",
			isErr:    assert.NoError,
		},
		{
			desc: "given query params are preserved",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=encoding('UTF-8')"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=encoding('UTF-8')",
			isErr:    assert.NoError,
		},
		{
			desc: "file:// prefix",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file://path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "file://path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
			isErr:    assert.NoError,
		},
		{
			desc: "file: prefix",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "file:/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
			isErr:    assert.NoError,
		},
		{
			desc: "windows file path gets normalized",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "C:\\path\\to\\a\\file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "C:/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
			isErr:    assert.NoError,
		},
		// Test cases from official SQLite docs: https://sqlite.org/c3ref/open.html#:~:text=URI%20filename%20examples
		// These assert current DSN behavior only, without SQLite-specific URI validation.
		{
			desc: "file:data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:data.db"
				return s
			}(),
			expected: "file:data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file:/home/fred/data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:/home/fred/data.db"
				return s
			}(),
			expected: "file:/home/fred/data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file:///home/fred/data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:///home/fred/data.db"
				return s
			}(),
			expected: "file:///home/fred/data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file://localhost/home/fred/data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file://localhost/home/fred/data.db"
				return s
			}(),
			expected: "file://localhost/home/fred/data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file://darkstar/home/fred/data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file://darkstar/home/fred/data.db"
				return s
			}(),
			expected: "file://darkstar/home/fred/data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file:///C:/Documents%20and%20Settings/fred/Desktop/data.db",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:///C:/Documents%20and%20Settings/fred/Desktop/data.db"
				return s
			}(),
			expected: "file:///C:/Documents%20and%20Settings/fred/Desktop/data.db",
			isErr:    assert.NoError,
		},
		{
			desc: "file:data.db?mode=ro&cache=private",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:data.db?mode=ro&cache=private"
				return s
			}(),
			expected: "file:data.db?mode=ro&cache=private",
			isErr:    assert.NoError,
		},
		{
			desc: "file:/home/fred/data.db?vfs=unix-dotfile",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:/home/fred/data.db?vfs=unix-dotfile"
				return s
			}(),
			expected: "file:/home/fred/data.db?vfs=unix-dotfile",
			isErr:    assert.NoError,
		},
		{
			desc: "file:data.db?mode=readonly",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:data.db?mode=readonly"
				return s
			}(),
			expected: "file:data.db?mode=readonly",
			isErr:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			db := NewSQLite(tt.settings)

			actual, err := db.DSN()
			tt.isErr(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
