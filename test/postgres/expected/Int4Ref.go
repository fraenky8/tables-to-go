package dto

import (
	"database/sql"
)

type Int4Ref struct {
	Int4Ref sql.NullInt64 `db:"int4_ref"`
}
