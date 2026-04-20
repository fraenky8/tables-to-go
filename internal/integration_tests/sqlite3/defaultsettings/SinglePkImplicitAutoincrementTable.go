package dto

import (
	"database/sql"
)

type SinglePkImplicitAutoincrementTable struct {
	Pk   int            `db:"pk"`
	Name sql.NullString `db:"name"`
}
