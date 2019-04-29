package tagger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/config"
	"github.com/fraenky8/tables-to-go/pkg/database"
)

func TestTaggers_GenerateTags(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *config.Settings
		column   database.Column
		expected string
	}{
		{
			desc: "enabled db-tag (default) without other tags generates only db-tags",
			settings: func() *config.Settings {
				s := config.NewSettings()
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
			settings: func() *config.Settings {
				s := config.NewSettings()
				s.TagsNoDb = true
				return s
			},
			column:   database.Column{},
			expected: "",
		},
		{
			desc: "default db-tag with enabled Mastermind-tag creates db- and Mastermind-tags",
			settings: func() *config.Settings {
				s := config.NewSettings()
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
			settings: func() *config.Settings {
				s := config.NewSettings()
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
			settings: func() *config.Settings {
				s := config.NewSettings()
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
			settings := test.settings()
			taggers := NewTaggers(settings)
			db := database.New(settings)
			actual := taggers.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
