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
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB},
			},
		},
		{
			desc: "explicit tags keep default db unless disabled",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"structable"}
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB, TagStructable},
			},
		},
		{
			desc: "explicit tags normalize aliases trims and dedupe",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{" db ", "stbl", "json", "sqlx", "JSON"}
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB, TagStructable, "json"},
			},
		},
		{
			desc: "empty and whitespace tags are ignored",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"", "   ", "\t", "structable"}
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB, TagStructable},
			},
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
			desc: "legacy tags structable adds structable",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructable = true
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB, TagStructable},
			},
		},
		{
			desc: "legacy tags structable only forces structable",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructableOnly = true
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagStructable},
			},
		},
		{
			desc: "legacy tags structable only overrides explicit custom tags",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"json", "structable"}
				s.TagsMastermindStructableOnly = true
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagStructable},
			},
		},
		{
			desc: "known and unknown tags are preserved in order",
			settings: func() *Settings {
				s := New()
				s.Tags = StringsFlag{"structable", "json", "yaml"}
				return s
			},
			expected: ResolvedTags{
				Tags: StringsFlag{TagDB, TagStructable, "json", "yaml"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.settings().ResolveTags()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSettings_UsesLegacyTagFlags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		settings func() *Settings
		expected bool
	}{
		{
			desc: "no legacy tags flags in use",
			settings: func() *Settings {
				return New()
			},
			expected: false,
		},
		{
			desc: "legacy tags no db in use",
			settings: func() *Settings {
				s := New()
				s.TagsNoDb = true
				return s
			},
			expected: true,
		},
		{
			desc: "legacy tags structable in use",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructable = true
				return s
			},
			expected: true,
		},
		{
			desc: "legacy tags structable only in use",
			settings: func() *Settings {
				s := New()
				s.TagsMastermindStructableOnly = true
				return s
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.settings().UsesLegacyTagFlags()
			assert.Equal(t, test.expected, actual)
		})
	}
}
