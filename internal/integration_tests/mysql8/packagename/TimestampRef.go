package models

import (
	"database/sql"
)

type TimestampRef struct {
	TimestampRef sql.NullTime `db:"timestamp_ref"`
}
