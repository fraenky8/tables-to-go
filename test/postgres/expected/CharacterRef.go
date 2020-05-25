package dto

import (
	"database/sql"
)

type CharacterRef struct {
	CharacterRef sql.NullString `db:"character_ref"`
}
