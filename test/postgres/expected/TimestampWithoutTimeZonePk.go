package dto

import (
	"time"
)

type TimestampWithoutTimeZonePk struct {
	TimestampWithoutTimeZonePk time.Time `db:"timestamp_without_time_zone_pk"`
}
