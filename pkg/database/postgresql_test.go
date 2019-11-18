package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func TestPostgresql_DSN(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *settings.Settings
		expected func(*settings.Settings) string
	}{
		{
			desc: "no username given, defaults to `postgres`",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypePostgresql
				return s
			},
			expected: func(s *settings.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "postgres", s.DbName, s.Pswd)
			},
		},
		{
			desc: "with given username, default gets overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypePostgresql
				s.User = "my_custom_user"
				return s
			},
			expected: func(s *settings.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "my_custom_user", s.DbName, s.Pswd)
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
