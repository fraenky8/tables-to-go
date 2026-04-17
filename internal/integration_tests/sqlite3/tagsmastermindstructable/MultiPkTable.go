package dto

import (
	"database/sql"
)

type MultiPkTable struct {
	PkA  int            `db:"pk_a" stbl:"pk_a,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	PkB  int            `db:"pk_b" stbl:"pk_b,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	Name sql.NullString `db:"name" stbl:"name"`
}
