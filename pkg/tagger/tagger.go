package tagger

import (
	"github.com/fraenky8/tables-to-go/pkg/database"
)

// Tagger interface for types of struct-tages
type Tagger interface {
	GenerateTag(db database.Database, column database.Column) string
}
