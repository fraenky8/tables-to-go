package dto

import (
	"time"
)

type TimeWithTimeZoneCheckPk struct {
	TimeWithTimeZoneCheckPk time.Time `db:"time_with_time_zone_check_pk"`
}
