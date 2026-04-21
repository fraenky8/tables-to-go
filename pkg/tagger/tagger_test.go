package tagger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

func TestTaggers_GenerateTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings *settings.Settings
		column   database.Column
		expected string
	}{
		{
			desc:     "enabled db-tag (default) without other tags generates only db-tags",
			settings: settings.New(),
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
			}(),
			column:   database.Column{},
			expected: "",
		},
		{
			desc: "explicit structable tag creates db and structable tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.Tags = settings.StringsFlag{"structable"}
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "unknown tag creates passthrough tag",
			settings: func() *settings.Settings {
				s := settings.New()
				s.Tags = settings.StringsFlag{"json"}
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" json:\"column_name\"`",
		},
		{
			desc: "mixed tags preserve configured order",
			settings: func() *settings.Settings {
				s := settings.New()
				s.Tags = settings.StringsFlag{"json", "structable"}
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" json:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "mixed new and legacy structable adds structable once",
			settings: func() *settings.Settings {
				s := settings.New()
				s.Tags = settings.StringsFlag{"db", "json"}
				s.TagsMastermindStructable = true
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" json:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "mixed new and legacy structable only keeps standalone structable",
			settings: func() *settings.Settings {
				s := settings.New()
				s.Tags = settings.StringsFlag{"db", "json"}
				s.TagsMastermindStructableOnly = true
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "disabled db-tag with enabled Mastermind-tag creates only Mastermind-tags",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "default db-tag with enabled standalone Mastermind-tag creates only standalone Mastermind-tag",
			settings: func() *settings.Settings {
				s := settings.New()
				s.TagsMastermindStructableOnly = true
				return s
			}(),
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings
			taggers := NewTaggers(s)
			db := database.New(s)
			actual := taggers.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
