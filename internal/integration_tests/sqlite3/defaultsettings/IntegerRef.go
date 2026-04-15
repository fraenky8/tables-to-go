package dto

import (
	"database/sql"
)

type IntegerRef struct {
	IntegerRef sql.NullString `db:"integer_ref"`
}
