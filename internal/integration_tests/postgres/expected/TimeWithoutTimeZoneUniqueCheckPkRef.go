package dto

import (
	"time"
)

type TimeWithoutTimeZoneUniqueCheckPkRef struct {
	TimeWithoutTimeZoneUniqueCheckPkRef time.Time `db:"time_without_time_zone_unique_check_pk_ref"`
}
