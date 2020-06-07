package dto

import (
	"time"
)

type TimestampWithTimeZonePkRef struct {
	TimestampWithTimeZonePkRef time.Time `db:"timestamp_with_time_zone_pk_ref"`
}
