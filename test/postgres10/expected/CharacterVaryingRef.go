package dto

import (
	"database/sql"
)

type CharacterVaryingRef struct {
	CharacterVaryingRef sql.NullString `db:"character_varying_ref"`
}
