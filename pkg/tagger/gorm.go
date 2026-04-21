package tagger

import (
	"strings"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
)

var (
	// GORM tag options are semicolon-separated and parser-special characters may
	// be escaped with backslashes (see gorm docs: Fields Tags). We escape
	// semicolons to keep values intact and quotes for valid Go struct tags.
	gormTagValueEscaper = strings.NewReplacer(
		`"`, `\"`,
		`;`, `\;`,
	)
)

// Gorm represents a gorm tag.
type Gorm struct{}

// GenerateTag for Gorm to satisfy the Tagger interface.
func (t Gorm) GenerateTag(db database.Database, column database.Column) string {
	parts := make([]string, 0, 5)

	parts = append(parts, "column:"+escapeGormTagValue(column.Name))

	if db.IsPrimaryKey(column) {
		parts = append(parts, "primaryKey")
	}

	if db.IsAutoIncrement(column) {
		parts = append(parts, "autoIncrement")
	}

	if !db.IsNullable(column) {
		parts = append(parts, "not null")
	}

	if column.DefaultValue.Valid {
		parts = append(parts, "default:"+escapeGormTagValue(column.DefaultValue.String))
	}

	return `gorm:"` + strings.Join(parts, ";") + `"`
}

func escapeGormTagValue(value string) string {
	return gormTagValueEscaper.Replace(value)
}
