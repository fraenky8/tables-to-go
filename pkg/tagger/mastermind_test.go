package tagger

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func TestMastermind_GenerateTag(t *testing.T) {
	type test struct {
		desc     string
		settings func() *settings.Settings
		column   database.Column
		expected string
	}

	tests := map[settings.DbType][]test{
		settings.DbTypePostgresql: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name: "column_name",
					ConstraintType: sql.NullString{
						String: "PRIMARY KEY",
						Valid:  true,
					},
				},
				expected: `stbl:"column_name,PRIMARY_KEY"`,
			},
			{
				desc: "PK and AI column generates Mastermind-tag with PK and AI indicator",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name: "column_name",
					ConstraintType: sql.NullString{
						String: "PRIMARY KEY",
						Valid:  true,
					},
					DefaultValue: sql.NullString{
						String: "nextval",
						Valid:  true,
					},
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
		settings.DbTypeMySQL: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name:      "column_name",
					ColumnKey: "PRI",
				},
				expected: `stbl:"column_name,PRIMARY_KEY"`,
			},
			{
				desc: "PK and AI column generates Mastermind-tag with PK and AI indicator",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name:      "column_name",
					ColumnKey: "PRI",
					Extra:     "auto_increment",
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
		settings.DbTypeSQLite: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypeSQLite
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator and AI indicator",
				settings: func() *settings.Settings {
					s := settings.New()
					s.DbType = settings.DbTypeSQLite
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: database.Column{
					Name:      "column_name",
					ColumnKey: "PK",
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
	}

	tagger := new(Mastermind)

	for dbType := range settings.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {
			tests := tests[dbType]
			for _, test := range tests {
				t.Run(test.desc, func(t *testing.T) {
					db := database.New(test.settings())
					actual := tagger.GenerateTag(db, test.column)
					assert.Equal(t, test.expected, actual)
				})
			}
		})
	}
}
