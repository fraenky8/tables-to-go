package models

import (
	"database/sql"
)

type BigintRef struct {
	BigintRef sql.NullInt64 `db:"bigint_ref"`
}
