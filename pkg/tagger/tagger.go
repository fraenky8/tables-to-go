package tagger

import (
	"strconv"
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
func NewTaggers(tags settings.ResolvedTags) *Taggers {
	t := &Taggers{
		taggers: make([]Tagger, 0, len(tags)),
	}
	for _, tag := range tags {
		switch tag {
		case settings.TagDB:
			t.taggers = append(t.taggers, Db{})
		case settings.TagGorm:
			t.taggers = append(t.taggers, Gorm{})
		case settings.TagStructable:
			t.taggers = append(t.taggers, Mastermind{})
		default:
			t.taggers = append(t.taggers, NewGeneric(tag))
		}
	}

	return t
}

// GenerateTag creates based on the enabled tags and the given database and column
// the tag for the struct field.
func (t *Taggers) GenerateTag(db database.Database, column database.Column) string {
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

	tags := sb.String()

	if len(tags) > 0 {
		tags = toStructTagLiteral(strings.TrimSpace(tags))
	}

	return tags
}

func toStructTagLiteral(tags string) string {
	if strings.ContainsRune(tags, '`') {
		return strconv.Quote(tags)
	}

	return "`" + tags + "`"
}
