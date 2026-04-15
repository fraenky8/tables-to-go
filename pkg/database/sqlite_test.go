package database

import (
	"net/url"
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
	}{
		{
			desc: "an invalid URL query param is ignored and returns DB name as is",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=\x00"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=\x00",
		},
		{
			desc: "plain db file without query params, default _pragma gets added",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=busy_timeout(5000)&_pragma=cache_size(20480)",
		},
		{
			desc: "given query params are preserved",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=encoding('UTF-8')"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=encoding('UTF-8')&_pragma=busy_timeout(5000)&_pragma=cache_size(20480)",
		},
		{
			desc: "busy_timeout is overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=busy_timeout(10000)"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(20480)",
		},
		{
			desc: "cache_size is overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=cache_size(10000)"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=cache_size(10000)&_pragma=busy_timeout(5000)",
		},
		{
			desc: "busy_timeout and cache_size are overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
		},
		{
			desc: "file:// prefix",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file://path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
		},
		{
			desc: "file: prefix",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "file:/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
		},
		{
			desc: "windows file path gets normalized",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				s.DbName = "C:\\path\\to\\a\\file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)"
				return s
			}(),
			expected: "/path/to/a/file.db?_pragma=busy_timeout(10000)&_pragma=cache_size(10000)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			db := NewSQLite(tt.settings)

			actual, err := url.QueryUnescape(db.DSN())
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
