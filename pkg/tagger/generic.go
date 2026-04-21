package tagger

import (
	"github.com/fraenky8/tables-to-go/v2/pkg/database"
)

// Generic is a passthrough tagger for arbitrary tag names.
type Generic struct {
	name string
}

// NewGeneric creates a Generic tagger.
func NewGeneric(name string) Generic {
	return Generic{name: name}
}

// GenerateTag satisfies the Tagger interface.
func (t Generic) GenerateTag(_ database.Database, column database.Column) string {
	if t.name == "" {
		return ""
	}

	return t.name + `:"` + column.Name + `"`
}
