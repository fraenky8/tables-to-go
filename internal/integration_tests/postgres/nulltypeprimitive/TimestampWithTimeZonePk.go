package dto

import (
	"time"
)

type TimestampWithTimeZonePk struct {
	TimestampWithTimeZonePk time.Time `db:"timestamp_with_time_zone_pk"`
}
