package tagger

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

func TestTaggers_GenerateTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		tags     settings.ResolvedTags
		column   database.Column
		expected string
	}{
		{
			desc: "enabled db-tag (default) without other tags generates only db-tags",
			tags: settings.ResolvedTags{
				settings.TagDB,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\"`",
		},
		{
			desc:     "disabled db-tag without other tags generates no tags",
			tags:     settings.ResolvedTags{},
			column:   database.Column{},
			expected: "",
		},
		{
			desc: "structable tag maps to mastermind tagger",
			tags: settings.ResolvedTags{
				settings.TagStructable,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "explicit structable tag creates db and structable tags",
			tags: settings.ResolvedTags{
				settings.TagDB,
				settings.TagStructable,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "unknown tag creates passthrough tag",
			tags: settings.ResolvedTags{
				settings.TagDB,
				"json",
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" json:\"column_name\"`",
		},
		{
			desc: "mixed tags preserve configured order",
			tags: settings.ResolvedTags{
				settings.TagDB,
				settings.TagGorm,
				"json",
				settings.TagStructable,
			},
			column: database.Column{
				Name:       "column_name",
				IsNullable: "YES",
			},
			expected: "`db:\"column_name\" gorm:\"column:column_name\" json:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "gorm tag maps to gorm tagger",
			tags: settings.ResolvedTags{
				settings.TagGorm,
			},
			column: database.Column{
				Name:       "column_name",
				IsNullable: "YES",
			},
			expected: "`gorm:\"column:column_name\"`",
		},
		{
			desc: "gorm tag with backtick in value uses interpreted string literal",
			tags: settings.ResolvedTags{
				settings.TagGorm,
			},
			column: database.Column{
				Name:       "column_name",
				IsNullable: "YES",
				DefaultValue: sql.NullString{
					String: "contains `tick",
					Valid:  true,
				},
			},
			expected: "\"gorm:\\\"column:column_name;default:contains `tick\\\"\"",
		},
		{
			desc: "gorm tag in combination keeps configured order",
			tags: settings.ResolvedTags{
				settings.TagDB,
				settings.TagGorm,
				"json",
			},
			column: database.Column{
				Name:       "column_name",
				IsNullable: "YES",
			},
			expected: "`db:\"column_name\" gorm:\"column:column_name\" json:\"column_name\"`",
		},
		{
			desc: "mixed new and legacy structable adds structable once",
			tags: settings.ResolvedTags{
				settings.TagDB,
				settings.TagStructable,
				"json",
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" stbl:\"column_name\" json:\"column_name\"`",
		},
		{
			desc: "mixed new and legacy structable only keeps standalone structable",
			tags: settings.ResolvedTags{
				settings.TagStructable,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "disabled db-tag with enabled Mastermind-tag creates only Mastermind-tags",
			tags: settings.ResolvedTags{
				settings.TagStructable,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "default db-tag with enabled standalone Mastermind-tag creates only standalone Mastermind-tag",
			tags: settings.ResolvedTags{
				settings.TagStructable,
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "empty tag gets ignored",
			tags: settings.ResolvedTags{
				settings.TagDB,
				"",
			},
			column: database.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\"`",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			taggers := NewTaggers(test.tags)
			db := database.New(settings.New())
			actual := taggers.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
