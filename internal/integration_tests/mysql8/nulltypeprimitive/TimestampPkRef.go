package dto

import (
	"time"
)

type TimestampPkRef struct {
	TimestampPkRef time.Time `db:"timestamp_pk_ref"`
}
