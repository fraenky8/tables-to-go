package dto

import (
	"time"
)

type TimeWithTimeZoneUniqueCheckPk struct {
	TimeWithTimeZoneUniqueCheckPk time.Time `db:"time_with_time_zone_unique_check_pk"`
}
