package dto

import (
	"database/sql"
)

type VarcharRef struct {
	VarcharRef sql.NullString `db:"varchar_ref"`
}
