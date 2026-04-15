package dto

import (
	"database/sql"
)

type RealRef struct {
	RealRef sql.NullString `db:"real_ref"`
}
