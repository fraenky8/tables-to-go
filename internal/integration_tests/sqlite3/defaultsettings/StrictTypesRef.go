package dto

import (
	"database/sql"
)

type StrictTypesRef struct {
	StrictTypesRef sql.NullInt64 `db:"strict_types_ref"`
}
