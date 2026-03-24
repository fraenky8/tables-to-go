package models

import (
	"database/sql"
)

type TimeRef struct {
	TimeRef sql.NullTime `db:"time_ref"`
}
