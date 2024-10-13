package dto

import (
	"database/sql"
)

type Int2Ref struct {
	Int2Ref sql.NullInt64 `db:"int2_ref"`
}
