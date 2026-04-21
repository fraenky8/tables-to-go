package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettings_ResolveTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings func() *Settings
		expected ResolvedTags
	}{
		{
			desc: "default settings resolve to db tag",
			settings: func() *Settings {
				return New()
			},
			expected: ResolvedTags{TagDB},
		},
		{
			desc: "explicit tags keep default db unless disabled",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"structable"}
				return s
			},
			expected: ResolvedTags{TagDB, TagStructable},
		},
		{
			desc: "explicit tags normalize known tags and preserve custom tag casing",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{" db ", "stbl", "json", "sqlx", "JSON", "STRUCTABLE"}
				return s
			},
			expected: ResolvedTags{TagDB, TagStructable, "json", "JSON"},
		},
		{
			desc: "empty and whitespace tags are ignored",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"", "   ", "\t", "structable"}
				return s
			},
			expected: ResolvedTags{TagDB, TagStructable},
		},
		{
			desc: "legacy tags no db removes default db",
			settings: func() *Settings {
				s := New()
				s.TagsNoDb = true
				return s
			},
			expected: ResolvedTags{},
		},
		{
			desc: "legacy tags no db removes explicit db tag",
			settings: func() *Settings {
				s := New()
				s.TagsNoDb = true
				s.Tags = StringsFlag{"db", "json"}
				return s
			},
			expected: ResolvedTags{"json"},
		},
		{
			desc: "legacy tags no db removes explicit sqlx alias",
			settings: func() *Settings {
				s := New()
				s.TagsNoDb = true
				s.Tags = StringsFlag{"sqlx", "json"}
				return s
			},
			expected: ResolvedTags{"json"},
		},
		{
			desc: "legacy tags structable adds structable",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructable = true
				return s
			},
			expected: ResolvedTags{TagDB, TagStructable},
		},
		{
			desc: "legacy tags structable only forces structable",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructableOnly = true
				return s
			},
			expected: ResolvedTags{TagStructable},
		},
		{
			desc: "legacy tags structable only overrides explicit custom tags",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"json", "structable"}
				s.TagsMastermindStructableOnly = true
				return s
			},
			expected: ResolvedTags{TagStructable},
		},
		{
			desc: "known and unknown tags are preserved in order",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"structable", "json", "yaml"}
				return s
			},
			expected: ResolvedTags{TagDB, TagStructable, "json", "yaml"},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.settings().ResolveTags()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestResolvedTags_removeTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc      string
		resolved  ResolvedTags
		removeTag string
		expected  ResolvedTags
	}{
		{
			desc:      "removing from empty tags keeps empty result",
			resolved:  ResolvedTags{},
			removeTag: TagDB,
			expected:  ResolvedTags{},
		},
		{
			desc:      "removing missing tag keeps tags unchanged",
			resolved:  ResolvedTags{TagDB, TagStructable, "json"},
			removeTag: "yaml",
			expected:  ResolvedTags{TagDB, TagStructable, "json"},
		},
		{
			desc:      "removing existing tag deletes it",
			resolved:  ResolvedTags{TagDB, TagStructable, "json"},
			removeTag: TagStructable,
			expected:  ResolvedTags{TagDB, "json"},
		},
		{
			desc:      "removing only tag leaves empty result",
			resolved:  ResolvedTags{TagDB},
			removeTag: TagDB,
			expected:  ResolvedTags{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.resolved
			actual.removeTag(test.removeTag)
			assert.Equal(t, test.expected, actual)
		})
	}
}
