package dto

import (
	"database/sql"
)

type SmallintRef struct {
	SmallintRef sql.NullInt64 `db:"smallint_ref"`
}
