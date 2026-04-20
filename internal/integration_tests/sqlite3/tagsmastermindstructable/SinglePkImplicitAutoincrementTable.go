package dto

import (
	"database/sql"
)

type SinglePkImplicitAutoincrementTable struct {
	Pk   int            `db:"pk" stbl:"pk,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	Name sql.NullString `db:"name" stbl:"name"`
}
