package tagger

import "github.com/fraenky8/tables-to-go/pkg"

// Mastermind represents the Masterminds/structable "stbl"-tag
type Mastermind string

// GenerateTag for Mastermind to satisfy the Tagger interface
func (t *Mastermind) GenerateTag(db pkg.Database, column pkg.Column) string {

	isPk := ""
	if db.IsPrimaryKey(column) {
		isPk = ",PRIMARY_KEY"
	}

	isAutoIncrement := ""
	if db.IsAutoIncrement(column) {
		isAutoIncrement = ",SERIAL,AUTO_INCREMENT"
	}

	return `stbl:"` + column.Name + isPk + isAutoIncrement + `"`
}
