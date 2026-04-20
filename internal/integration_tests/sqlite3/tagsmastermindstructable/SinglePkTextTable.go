package dto

import (
	"database/sql"
)

type SinglePkTextTable struct {
	Pk   string         `db:"pk" stbl:"pk,PRIMARY_KEY"`
	Name sql.NullString `db:"name" stbl:"name"`
}
