package dto

import (
	"database/sql"
)

type DatetimeRef struct {
	DatetimeRef sql.NullTime `db:"datetime_ref"`
}
