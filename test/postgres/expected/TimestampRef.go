package dto

import (
	pg "github.com/lib/pq"
)

type TimestampRef struct {
	TimestampRef pg.NullTime `db:"timestamp_ref"`
}
