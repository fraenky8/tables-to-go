package dto

import (
	"time"
)

type TimestampWithoutTimeZoneNnUniqueCheckPk struct {
	TimestampWithoutTimeZoneNnUniqueCheckPk time.Time `db:"timestamp_without_time_zone_nn_unique_check_pk"`
}
