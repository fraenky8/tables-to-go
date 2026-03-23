package dto

import (
	"time"
)

type TimestampRef struct {
	TimestampRef time.Time `db:"timestamp_ref"`
}
