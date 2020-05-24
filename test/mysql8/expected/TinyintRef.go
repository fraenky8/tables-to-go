package dto

import (
	"database/sql"
)

type TinyintRef struct {
	TinyintRef sql.NullInt64 `db:"tinyint_ref"`
}
