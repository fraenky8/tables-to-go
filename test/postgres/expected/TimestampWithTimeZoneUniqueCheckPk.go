package dto

import (
	"time"
)

type TimestampWithTimeZoneUniqueCheckPk struct {
	TimestampWithTimeZoneUniqueCheckPk time.Time `db:"timestamp_with_time_zone_unique_check_pk"`
}
