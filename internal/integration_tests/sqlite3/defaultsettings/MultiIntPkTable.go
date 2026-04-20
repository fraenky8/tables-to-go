package dto

import (
	"database/sql"
)

type MultiIntPkTable struct {
	PkA  int            `db:"pk_a"`
	PkB  int            `db:"pk_b"`
	Name sql.NullString `db:"name"`
}
