package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

func TestPostgresql_DSN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings func() *settings.Settings
		expected func(*settings.Settings) string
	}{
		{
			desc: "no username given, defaults to `postgres`",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				return s
			},
			expected: func(s *settings.Settings) string {
				expected := "postgres://postgres:"
				expected += s.Pswd + "@" + s.Host + ":" + s.Port + "/"
				expected += s.DbName + "?sslmode=" + s.SSLMode
				return expected
			},
		},
		{
			desc: "with given username, default gets overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "my_custom_user"
				return s
			},
			expected: func(s *settings.Settings) string {
				expected := "postgres://my_custom_user:"
				expected += s.Pswd + "@" + s.Host + ":" + s.Port + "/"
				expected += s.DbName + "?sslmode=" + s.SSLMode
				return expected
			},
		},
		{
			desc: "with given username and socket, default gets overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "my_custom_user"
				s.Pswd = "mysecretpassword"
				s.Socket = "/tmp"
				return s
			},
			expected: func(s *settings.Settings) string {
				expected := "postgres://my_custom_user:mysecretpassword@"
				expected += "?" + s.Socket + "&" + s.Port + "&sslmode=" + s.SSLMode
				return expected
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := NewPostgresql(test.settings())
			actual := db.DSN()
			assert.Equal(t, test.expected(db.Settings), actual)
		})
	}
}

func TestPostgresql_andInClause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc         string
		field        string
		params       []string
		args         []any
		expected     string
		expectedArgs []any
	}{
		{
			desc:         "empty field returns early",
			params:       nil,
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "nil params returns early",
			field:        "table_name",
			params:       nil,
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "zero params returns early",
			field:        "table_name",
			params:       []string{},
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "one param returns AND IN clause",
			field:        "table_name",
			params:       []string{"foo"},
			args:         []any{},
			expected:     "AND table_name IN ($1)",
			expectedArgs: []any{"foo"},
		},
		{
			desc:         "multiple params returns AND IN clause",
			field:        "table_name",
			params:       []string{"foo", "bar", "baz", "qux"},
			args:         []any{},
			expected:     "AND table_name IN ($1,$2,$3,$4)",
			expectedArgs: []any{"foo", "bar", "baz", "qux"},
		},
		{
			desc:         "multiple params and existing args returns AND IN clause",
			field:        "table_name",
			params:       []string{"baz", "qux", "quux", "corge"},
			args:         []any{"foo", "bar"},
			expected:     "AND table_name IN ($3,$4,$5,$6)",
			expectedArgs: []any{"foo", "bar", "baz", "qux", "quux", "corge"},
		},
	}
	for i, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := (&Postgresql{}).andInClause(tt.field, tt.params, &tests[i].args)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedArgs, tests[i].args)
		})
	}
}
