package dto

import (
	"time"
)

type TimestampWithoutTimeZoneNnUniqueCheckPkRef struct {
	TimestampWithoutTimeZoneNnUniqueCheckPkRef time.Time `db:"timestamp_without_time_zone_nn_unique_check_pk_ref"`
}
