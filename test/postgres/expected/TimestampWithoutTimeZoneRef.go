package dto

import (
	pg "github.com/lib/pq"
)

type TimestampWithoutTimeZoneRef struct {
	TimestampWithoutTimeZoneRef pg.NullTime `db:"timestamp_without_time_zone_ref"`
}
