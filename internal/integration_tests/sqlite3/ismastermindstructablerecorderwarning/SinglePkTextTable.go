package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type SinglePkTextTable struct {
	structable.Recorder

	Pk   string         `db:"pk"`
	Name sql.NullString `db:"name"`
}
