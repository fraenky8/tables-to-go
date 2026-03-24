package models

import (
	"database/sql"
)

type MediumintRef struct {
	MediumintRef sql.NullInt64 `db:"mediumint_ref"`
}
