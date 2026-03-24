package models

import (
	"database/sql"
)

type YearRef struct {
	YearRef sql.NullTime `db:"year_ref"`
}
