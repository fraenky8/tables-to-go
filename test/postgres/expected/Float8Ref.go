package dto

import (
	"database/sql"
)

type Float8Ref struct {
	Float8Ref sql.NullFloat64 `db:"float8_ref"`
}
