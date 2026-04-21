package tagger

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

func TestGorm_GenerateTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings func() *settings.Settings
		column   database.Column
		expected string
	}{
		{
			desc:     "minimal field emits column option only",
			settings: settings.New,
			column: database.Column{
				Name:       "column_name",
				IsNullable: "YES",
			},
			expected: `gorm:"column:column_name"`,
		},
		{
			desc: "full metadata emits supported options in deterministic order",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeMySQL
				return s
			},
			column: database.Column{
				Name:       "id",
				IsNullable: "NO",
				DefaultValue: sql.NullString{
					String: "nextval('users_id_seq'::regclass)",
					Valid:  true,
				},
				ColumnKey: "PRI",
				Extra:     "auto_increment",
			},
			expected: `gorm:"column:id;primaryKey;autoIncrement;not null;default:nextval('users_id_seq'::regclass)"`,
		},
		{
			desc: "sqlite integer pk with autoincrement emits primaryKey and autoIncrement",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				return s
			},
			column: database.Column{
				Name:       "id",
				IsNullable: "YES",
				ColumnKey:  "PK",
				Extra:      "auto_increment",
			},
			expected: `gorm:"column:id;primaryKey;autoIncrement"`,
		},
		{
			desc: "sqlite pk without autoincrement emits primaryKey only",
			settings: func() *settings.Settings {
				s := settings.New()
				s.DbType = settings.DBTypeSQLite
				return s
			},
			column: database.Column{
				Name:       "code",
				IsNullable: "YES",
				ColumnKey:  "PK",
			},
			expected: `gorm:"column:code;primaryKey"`,
		},
		{
			desc:     "data type is intentionally omitted and inferred by gorm",
			settings: settings.New,
			column: database.Column{
				Name:       "nickname",
				DataType:   "character varying",
				IsNullable: "YES",
			},
			expected: `gorm:"column:nickname"`,
		},
		{
			desc:     "special characters are escaped for gorm tag parser and go tags",
			settings: settings.New,
			column: database.Column{
				Name:       `weird;name\path"quote`,
				IsNullable: "YES",
				DefaultValue: sql.NullString{
					String: `a;b\c"d`,
					Valid:  true,
				},
			},
			expected: `gorm:"column:weird\;name\path\"quote;default:a\;b\c\"d"`,
		},
		{
			desc:     "backslashes are preserved and not double-escaped",
			settings: settings.New,
			column: database.Column{
				Name:       `windows_path`,
				IsNullable: "YES",
				DefaultValue: sql.NullString{
					String: `C:\tmp\file.txt`,
					Valid:  true,
				},
			},
			expected: `gorm:"column:windows_path;default:C:\tmp\file.txt"`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := database.New(test.settings())
			actual := Gorm{}.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
