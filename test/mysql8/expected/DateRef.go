package dto

import (
	"database/sql"
)

type DateRef struct {
	DateRef sql.NullTime `db:"date_ref"`
}
