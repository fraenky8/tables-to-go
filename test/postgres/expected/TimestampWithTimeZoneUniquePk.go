package dto

import (
	"time"
)

type TimestampWithTimeZoneUniquePk struct {
	TimestampWithTimeZoneUniquePk time.Time `db:"timestamp_with_time_zone_unique_pk"`
}
