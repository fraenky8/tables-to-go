package dto

import (
	"time"
)

type TimestampWithoutTimeZoneDefConstUniqueCheckPk struct {
	TimestampWithoutTimeZoneDefConstUniqueCheckPk time.Time `db:"timestamp_without_time_zone_def_const_unique_check_pk"`
}
