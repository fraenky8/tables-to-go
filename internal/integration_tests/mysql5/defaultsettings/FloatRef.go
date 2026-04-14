package dto

import (
	"database/sql"
)

type FloatRef struct {
	FloatRef sql.NullFloat64 `db:"float_ref"`
}
