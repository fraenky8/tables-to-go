package dto

import (
	"database/sql"
)

type SinglePkTextTable struct {
	Pk   string         `db:"pk"`
	Name sql.NullString `db:"name"`
}
