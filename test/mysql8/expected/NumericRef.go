package dto

import (
	"database/sql"
)

type NumericRef struct {
	NumericRef sql.NullFloat64 `db:"numeric_ref"`
}
