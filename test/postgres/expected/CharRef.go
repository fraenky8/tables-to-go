package dto

import (
	"database/sql"
)

type CharRef struct {
	CharRef sql.NullString `db:"char_ref"`
}
