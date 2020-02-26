package dto

import (
	"time"
)

type TimestampWithoutTimeZoneUniquePk struct {
	TimestampWithoutTimeZoneUniquePk time.Time `db:"timestamp_without_time_zone_unique_pk"`
}
