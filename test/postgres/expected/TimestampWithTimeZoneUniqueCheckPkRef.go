package dto

import (
	"time"
)

type TimestampWithTimeZoneUniqueCheckPkRef struct {
	TimestampWithTimeZoneUniqueCheckPkRef time.Time `db:"timestamp_with_time_zone_unique_check_pk_ref"`
}
