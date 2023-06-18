package tagger

import (
	"github.com/fraenky8/tables-to-go/pkg/database"
)

// Db is the standard "db"-tag.
type Db struct{}

// GenerateTag for Db to satisfy the Tagger interface.
func (t Db) GenerateTag(_ database.Database, column database.Column) string {
	return `db:"` + column.Name + `"`
}
