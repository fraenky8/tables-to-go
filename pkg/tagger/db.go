package tagger

import "github.com/fraenky8/tables-to-go/pkg"

// Db is the standard "db"-tag
type Db string

// GenerateTag for Db to satisfy the Tagger interface
func (t *Db) GenerateTag(db pkg.Database, column pkg.Column) string {
	return `db:"` + column.Name + `"`
}
