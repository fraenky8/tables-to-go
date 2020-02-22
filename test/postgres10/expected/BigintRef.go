package dto

import (
	"database/sql"
)

type BigintRef struct {
	BigintRef sql.NullInt64 `db:"bigint_ref"`
}
