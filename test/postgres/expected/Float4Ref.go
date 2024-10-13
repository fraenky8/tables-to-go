package dto

import (
	"database/sql"
)

type Float4Ref struct {
	Float4Ref sql.NullFloat64 `db:"float4_ref"`
}
