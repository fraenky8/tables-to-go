package tagger

import "github.com/fraenky8/tables-to-go"

// Db is the standard "db"-tag
type Db string

// GenerateTag for Db to satisfy the Tagger interface
func (t *Db) GenerateTag(db tablestogo.Database, column tablestogo.Column) string {
	return `db:"` + column.Name + `"`
}
