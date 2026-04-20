package dto

import (
	"database/sql"
)

type MultiIntPkTable struct {
	PkA  int            `db:"pk_a" stbl:"pk_a,PRIMARY_KEY"`
	PkB  int            `db:"pk_b" stbl:"pk_b,PRIMARY_KEY"`
	Name sql.NullString `db:"name" stbl:"name"`
}
