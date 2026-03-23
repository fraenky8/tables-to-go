package dto

import (
	"time"
)

type TimestampWithoutTimeZoneCheckPk struct {
	TimestampWithoutTimeZoneCheckPk time.Time `db:"timestamp_without_time_zone_check_pk"`
}
