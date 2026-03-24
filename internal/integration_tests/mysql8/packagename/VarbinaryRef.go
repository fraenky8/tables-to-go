package models

import (
	"database/sql"
)

type VarbinaryRef struct {
	VarbinaryRef sql.NullString `db:"varbinary_ref"`
}
