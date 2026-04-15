package dto

import (
	"database/sql"
)

type StrictTypesRef struct {
	StrictTypesRef sql.NullString `db:"strict_types_ref"`
}
