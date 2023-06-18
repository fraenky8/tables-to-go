package dto

import (
	"database/sql"
)

type TimestampWithoutTimeZoneRef struct {
	TimestampWithoutTimeZoneRef sql.NullTime `db:"timestamp_without_time_zone_ref"`
}
