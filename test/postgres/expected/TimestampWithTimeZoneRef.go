package dto

import (
	"database/sql"
)

type TimestampWithTimeZoneRef struct {
	TimestampWithTimeZoneRef sql.NullTime `db:"timestamp_with_time_zone_ref"`
}
