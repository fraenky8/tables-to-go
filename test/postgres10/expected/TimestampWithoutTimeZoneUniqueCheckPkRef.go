package dto

import (
	"time"
)

type TimestampWithoutTimeZoneUniqueCheckPkRef struct {
	TimestampWithoutTimeZoneUniqueCheckPkRef time.Time `db:"timestamp_without_time_zone_unique_check_pk_ref"`
}
