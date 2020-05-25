package dto

import (
	"time"
)

type TimeWithoutTimeZoneUniqueCheckPk struct {
	TimeWithoutTimeZoneUniqueCheckPk time.Time `db:"time_without_time_zone_unique_check_pk"`
}
