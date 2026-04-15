package dto

import (
	"database/sql"
)

type RealRef struct {
	RealRef sql.NullFloat64 `db:"real_ref"`
}
