// +build !integration

package tagger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func TestTaggers_GenerateTags(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *settings.Settings
		column   database.Column
		expected string
	}{
		{
			desc: "enabled db-tag (default) without other tags generates only db-tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = false
				return s
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\"`",
		},
		{
			desc: "disabled db-tag without other tags generates no tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = true
				return s
			},
			column:   database.Column{},
			expected: "",
		},
		{
			desc: "default db-tag with enabled Mastermind-tag creates db- and Mastermind-tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = false
				s.TagsMastermindStructable = true
				return s
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "disabled db-tag with enabled Mastermind-tag creates only Mastermind-tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "default db-tag with enabled standalone Mastermind-tag creates only standalone Mastermind-tag",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = false
				s.TagsMastermindStructableOnly = true
				return s
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings()
			taggers := NewTaggers(s)
			db := database.New(s)
			actual := taggers.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
