package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func TestMySQL_DSN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings func() *settings.Settings
		expected func(*settings.Settings) string
	}{
		{
			desc: "no username given, defaults to `root` with tcp",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeMySQL
				s.Pswd = "mysecretpassword"
				s.DbName = "my-cool-db"
				s.Port = "3306"
				return s
			},
			expected: func(*settings.Settings) string {
				return "root:mysecretpassword@tcp(127.0.0.1:3306)/my-cool-db"
			},
		},
		{
			desc: "username given, with socket",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeMySQL
				s.User = "admin"
				s.Pswd = "mysecretpassword"
				s.DbName = "my-cool-db"
				s.Socket = "/tmp/mysql.sock"
				return s
			},
			expected: func(*settings.Settings) string {
				return "admin:mysecretpassword@unix(/tmp/mysql.sock)/my-cool-db"
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := NewMySQL(test.settings())
			actual := db.DSN()
			assert.Equal(t, test.expected(db.Settings), actual)
		})
	}
}
