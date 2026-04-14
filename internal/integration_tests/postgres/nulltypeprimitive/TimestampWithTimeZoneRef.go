package dto

import (
	"time"
)

type TimestampWithTimeZoneRef struct {
	TimestampWithTimeZoneRef *time.Time `db:"timestamp_with_time_zone_ref"`
}
