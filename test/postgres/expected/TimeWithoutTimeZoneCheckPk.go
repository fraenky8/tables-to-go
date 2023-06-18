package dto

import (
	"time"
)

type TimeWithoutTimeZoneCheckPk struct {
	TimeWithoutTimeZoneCheckPk time.Time `db:"time_without_time_zone_check_pk"`
}
