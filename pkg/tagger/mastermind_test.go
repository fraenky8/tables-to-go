package tagger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/settings"
)

type mastermindMockDb struct {
	mock.Mock
	database.Database
}

func newMastermindMockDb(db database.Database) *mastermindMockDb {
	return &mastermindMockDb{Database: db}
}

func (db *mastermindMockDb) IsPrimaryKey(column database.Column) bool {
	db.Called(column)
	return column.ColumnKey == "PK"
}

func (db *mastermindMockDb) IsAutoIncrement(column database.Column) bool {
	db.Called(column)
	return column.Extra == "AI"
}

func TestMastermind_GenerateTag(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *settings.Settings
		database func(settings *settings.Settings, column database.Column) database.Database
		column   database.Column
		expected string
	}{
		{
			desc: "non PK column generates standard Mastermind-tag",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			},
			database: func(settings *settings.Settings, column database.Column) database.Database {
				db := database.New(settings)
				mdb := newMastermindMockDb(db)
				mdb.
					On("IsPrimaryKey", column).
					Return(false).
					On("IsAutoIncrement", column).
					Return(false)
				return mdb
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
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			},
			database: func(settings *settings.Settings, column database.Column) database.Database {
				db := database.New(settings)
				mdb := newMastermindMockDb(db)
				mdb.
					On("IsPrimaryKey", column).
					Return(true).
					On("IsAutoIncrement", column).
					Return(false)
				return mdb
			},
			column: database.Column{
				Name:      "column_name",
				ColumnKey: "PK",
			},
			expected: `stbl:"column_name,PRIMARY_KEY"`,
		},
		{
			desc: "PK and AI column generates Mastermind-tag with PK and AI indicator",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			},
			database: func(settings *settings.Settings, column database.Column) database.Database {
				db := database.New(settings)
				mdb := newMastermindMockDb(db)
				mdb.
					On("IsPrimaryKey", column).
					Return(true).
					On("IsAutoIncrement", column).
					Return(true)
				return mdb
			},
			column: database.Column{
				Name:      "column_name",
				ColumnKey: "PK",
				Extra:     "AI",
			},
			expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
		},
	}

	tagger := new(Mastermind)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := test.database(test.settings(), test.column)
			actual := tagger.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
