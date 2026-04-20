package tagger

import (
	"strings"
	"sync"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

var stringPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

// Tagger interface for types of struct-tags.
type Tagger interface {
	GenerateTag(db database.Database, column database.Column) string
}

// Taggers represents the supported tags to generate.
type Taggers struct {
	taggers []Tagger
}

// NewTaggers is the constructor function to create the supported taggers.
func NewTaggers(s *settings.Settings) *Taggers {
	resolved := s.ResolveTags()

	t := &Taggers{taggers: make([]Tagger, 0, len(resolved.Tags))}
	for _, tag := range resolved.Tags {
		switch tag {
		case settings.TagDB:
			t.taggers = append(t.taggers, new(Db))
		case settings.TagStructable:
			t.taggers = append(t.taggers, new(Mastermind))
		default:
			t.taggers = append(t.taggers, NewGeneric(tag))
		}
	}

	return t
}

// GenerateTag creates based on the enabled tags and the given database and column
// the tag for the struct field.
func (t *Taggers) GenerateTag(db database.Database, column database.Column) (tags string) {
	sb := stringPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		stringPool.Put(sb)
	}()

	for i := range t.taggers {
		tag := t.taggers[i].GenerateTag(db, column)
		if tag == "" {
			continue
		}

		sb.WriteString(tag)
		sb.WriteString(" ")
	}

	tags = sb.String()

	if len(tags) > 0 {
		tags = "`" + strings.TrimSpace(tags) + "`"
	}

	return tags
}
