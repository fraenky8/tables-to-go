package dto

import (
	"time"
)

type TimestampWithoutTimeZoneUniqueCheckPk struct {
	TimestampWithoutTimeZoneUniqueCheckPk time.Time `db:"timestamp_without_time_zone_unique_check_pk"`
}
