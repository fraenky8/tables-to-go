package database

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func TestSQLite_DSN(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *settings.Settings
		expected string
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc: "no username or password given, no authentication in DNS string",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db"
				return s
			},
			expected: "path/to/a/file.db",
			isError:  assert.NoError,
		},
		{
			desc: "with given username, authentication is enabled in DNS string",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db"
				s.User = "username"
				return s
			},
			expected: "path/to/a/file.db?_auth&_auth_user=username&_auth_pass=",
			isError:  assert.NoError,
		},
		{
			desc: "with given password, authentication is enabled in DNS string",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db"
				s.Pswd = "p4assw0rd"
				return s
			},
			expected: "path/to/a/file.db?_auth&_auth_user=&_auth_pass=p4assw0rd",
			isError:  assert.NoError,
		},
		{
			desc: "with given username and password, authentication is enabled in DNS string",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db"
				s.User = "username"
				s.Pswd = "p4assw0rd"
				return s
			},
			expected: "path/to/a/file.db?_auth&_auth_user=username&_auth_pass=p4assw0rd",
			isError:  assert.NoError,
		},
		{
			desc: "with existing username and password, authentication in DB name is overwritten",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db?_auth&_auth_user=username&_auth_pass=p4assw0rd"
				s.User = "new_username"
				s.Pswd = "new_p4assw0rd"
				return s
			},
			expected: "path/to/a/file.db?_auth&_auth_user=new_username&_auth_pass=new_p4assw0rd",
			isError:  assert.NoError,
		},
		{
			desc: "with existing username and password and additional option at the end, " +
				"authentication in DB name is overwritten and options are preserved",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db?_auth&_auth_user=username&_auth_pass=p4assw0rd&cache=shared"
				s.User = "new_username"
				s.Pswd = "new_p4assw0rd"
				return s
			},
			expected: "path/to/a/file.db?_auth&_auth_user=new_username&_auth_pass=new_p4assw0rd&cache=shared",
			isError:  assert.NoError,
		},
		{
			desc: "with existing username and password and additional option at the beginning, " +
				"authentication in DB name is overwritten and options are preserved",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = "path/to/a/file.db?cache=shared&_auth&_auth_user=username&_auth_pass=p4assw0rd"
				s.User = "new_username"
				s.Pswd = "new_p4assw0rd"
				return s
			},
			expected: "path/to/a/file.db?cache=shared&_auth&_auth_user=new_username&_auth_pass=new_p4assw0rd",
			isError:  assert.NoError,
		},
		{
			desc: "invalid dns returns raw dns string",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DbTypeSQLite
				s.DbName = ":123:456"
				s.User = "new_username"
				s.Pswd = "new_p4assw0rd"
				return s
			},
			expected: ":123:456",
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := NewSQLite(test.settings())
			dsn := db.DSN()

			actual, err := url.Parse(dsn)
			test.isError(t, err)
			if err != nil {
				assert.Equal(t, test.expected, dsn)
				return
			}

			expected, err := url.Parse(dsn)
			assert.NoError(t, err)

			assert.Equal(t, expected.Query(), actual.Query())
		})
	}
}
