package dto

import (
	"database/sql"
)

type TextRef struct {
	TextRef sql.NullString `db:"text_ref"`
}
