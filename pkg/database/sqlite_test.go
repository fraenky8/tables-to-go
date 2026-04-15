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
			desc: "plain db file without query params, default _pragma gets added",
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
