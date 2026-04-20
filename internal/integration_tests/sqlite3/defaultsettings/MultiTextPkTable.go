package dto

import (
	"database/sql"
)

type MultiTextPkTable struct {
	PkA  string         `db:"pk_a"`
	PkB  string         `db:"pk_b"`
	Name sql.NullString `db:"name"`
}
