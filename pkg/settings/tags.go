package settings

import (
	"slices"
	"strings"
)

// Special Tags
const (
	TagDB         = "db"
	TagStructable = "structable"
)

var (
	canonicalTags = map[string]string{
		TagDB:         TagDB,
		TagStructable: TagStructable,

		"sqlx": TagDB,
		"stbl": TagStructable,
	}
)

// ResolvedTags is the normalized tag configuration derived from new and legacy
// tag-related settings.
type ResolvedTags StringsFlag

// ResolveTags resolves tag configuration from the generic tags list and legacy
// tag booleans.
func (s *Settings) ResolveTags() ResolvedTags {
	if s.TagsMastermindStructableOnly {
		return ResolvedTags{TagStructable}
	}

	var resolved ResolvedTags
	// Backwards compatibility, always create `db` tags except s.TagsNoDb was provided.
	resolved.addTag(TagDB)

	if s.TagsMastermindStructable {
		resolved.addTag(TagStructable)
	}

	for i := range s.Tags {
		resolved.addTag(s.Tags[i])
	}

	// Yes, we remove the `db` tag again even if we added it previously to not
	// allow to reintroduce it via s.Tags db although s.TagsNoDb was provided.
	if s.TagsNoDb {
		resolved.removeTag(TagDB)
	}

	return resolved
}

func (r *ResolvedTags) addTag(value string) {
	tag := normalizeTag(value)
	if tag == "" {
		return
	}

	if slices.Contains(*r, tag) {
		return
	}
	*r = append(*r, tag)
}

func (r *ResolvedTags) removeTag(tag string) {
	i := slices.Index(*r, tag)
	if i < 0 {
		return
	}

	*r = append((*r)[:i], (*r)[i+1:]...)
}

func normalizeTag(value string) string {
	tag := strings.TrimSpace(value)
	if tag == "" {
		return ""
	}

	if t, ok := canonicalTags[strings.ToLower(tag)]; ok {
		return t
	}

	return tag
}
