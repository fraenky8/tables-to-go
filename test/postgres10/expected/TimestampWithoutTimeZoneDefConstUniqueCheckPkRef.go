package dto

import (
	"time"
)

type TimestampWithoutTimeZoneDefConstUniqueCheckPkRef struct {
	TimestampWithoutTimeZoneDefConstUniqueCheckPkRef time.Time `db:"timestamp_without_time_zone_def_const_unique_check_pk_ref"`
}
