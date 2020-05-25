package dto

import (
	"time"
)

type TimeWithoutTimeZoneDefConstUniqueCheckPk struct {
	TimeWithoutTimeZoneDefConstUniqueCheckPk time.Time `db:"time_without_time_zone_def_const_unique_check_pk"`
}
