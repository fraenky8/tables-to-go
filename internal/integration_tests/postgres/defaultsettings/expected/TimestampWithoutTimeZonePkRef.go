package dto

import (
	"time"
)

type TimestampWithoutTimeZonePkRef struct {
	TimestampWithoutTimeZonePkRef time.Time `db:"timestamp_without_time_zone_pk_ref"`
}
