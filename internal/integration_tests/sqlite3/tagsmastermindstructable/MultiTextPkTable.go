package dto

import (
	"database/sql"
)

type MultiTextPkTable struct {
	PkA  string         `db:"pk_a" stbl:"pk_a,PRIMARY_KEY"`
	PkB  string         `db:"pk_b" stbl:"pk_b,PRIMARY_KEY"`
	Name sql.NullString `db:"name" stbl:"name"`
}
