package dto

import (
	pg "github.com/lib/pq"
)

type TimestampWithTimeZoneRef struct {
	TimestampWithTimeZoneRef pg.NullTime `db:"timestamp_with_time_zone_ref"`
}
