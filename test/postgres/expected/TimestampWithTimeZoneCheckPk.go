package dto

import (
	"time"
)

type TimestampWithTimeZoneCheckPk struct {
	TimestampWithTimeZoneCheckPk time.Time `db:"timestamp_with_time_zone_check_pk"`
}
