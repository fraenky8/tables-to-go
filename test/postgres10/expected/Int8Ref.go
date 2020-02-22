package dto

import (
	"database/sql"
)

type Int8Ref struct {
	Int8Ref sql.NullInt64 `db:"int8_ref"`
}
