package dto

import (
	"database/sql"
)

type IntegerRef struct {
	IntegerRef sql.NullInt64 `db:"integer_ref"`
}
