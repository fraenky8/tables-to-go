package settings

import (
	"fmt"
	"reflect"
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

// ResolveTags computes and caches effective tags from Settings fields.
// Call once after Settings initialization and after (re)configuring tag-related
// fields, before Verify. Changes to tag-related fields made after this call
// are not reflected until ResolveTags is called again.
func (s *Settings) ResolveTags() {
	if s.TagsMastermindStructableOnly {
		s.tags = ResolvedTags{TagStructable}
		return
	}
	// Make this method stateless and resetting any resolved tags before so we
	// don't leak additional tags when invoked multiple times.
	s.tags = s.tags[:0]

	// Backwards compatibility, always create `db` tags except s.TagsNoDb was provided.
	s.tags.addTag(TagDB)

	if s.TagsMastermindStructable {
		s.tags.addTag(TagStructable)
	}

	for i := range s.Tags {
		s.tags.addTag(s.Tags[i])
	}

	// Yes, we remove the `db` tag again even if we added it previously to not
	// allow to reintroduce it via s.Tags db although s.TagsNoDb was provided.
	if s.TagsNoDb {
		s.tags.removeTag(TagDB)
	}
}

// Validate validates resolved tags.
func (r ResolvedTags) Validate() error {
	for i := range r {
		if err := validateTagKey(r[i]); err != nil {
			return err
		}
	}

	return nil
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

func validateTagKey(value string) error {
	tag := strings.TrimSpace(value)
	if tag == "" {
		return fmt.Errorf("invalid tag key %q: key must not be empty", value)
	}

	if strings.ContainsRune(tag, '`') {
		return fmt.Errorf("invalid tag key %q: key must not contain backtick", tag)
	}

	if _, ok := reflect.StructTag(tag + `:"x"`).Lookup(tag); !ok {
		return fmt.Errorf("invalid tag key %q: key must follow Go struct tag key syntax", tag)
	}

	return nil
}
