package dto

import (
	"database/sql"
)

type SinglePkExplicitAutoincrementTable struct {
	Pk   int            `db:"pk"`
	Name sql.NullString `db:"name"`
}
