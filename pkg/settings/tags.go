package settings

import (
	"slices"
	"strings"
)

const (
	TagDB         = "db"
	TagStructable = "structable"
)

var tagAliases = map[string]string{
	"sqlx": TagDB,
	"stbl": TagStructable,
}

// ResolvedTags is the normalized tag configuration derived from new and legacy
// tag-related settings.
type ResolvedTags struct {
	Tags StringsFlag
}

// ResolveTags resolves tag configuration from the generic tags list and legacy
// tag booleans.
func (s *Settings) ResolveTags() ResolvedTags {
	var resolved ResolvedTags

	if !s.TagsNoDb {
		resolved.appendTag(TagDB)
	}

	for i := range s.Tags {
		resolved.appendTag(s.Tags[i])
	}

	if s.TagsMastermindStructableOnly {
		resolved.Tags = StringsFlag{TagStructable}
		return resolved
	}

	if s.TagsMastermindStructable {
		resolved.appendTag(TagStructable)
	}

	return resolved
}

// UsesLegacyTagFlags returns true when any legacy tag compatibility flag is in use.
func (s *Settings) UsesLegacyTagFlags() bool {
	return s.TagsNoDb || s.TagsMastermindStructable || s.TagsMastermindStructableOnly
}

func (r *ResolvedTags) appendTag(value string) {
	tag := normalizeTag(value)
	if tag == "" {
		return
	}

	if slices.Contains(r.Tags, tag) {
		return
	}
	r.Tags = append(r.Tags, tag)
}

func normalizeTag(value string) string {
	tag := strings.ToLower(strings.TrimSpace(value))

	normalized, ok := tagAliases[tag]
	if ok {
		return normalized
	}

	return tag
}
