package dto

import (
	"time"
)

type TimeWithTimeZoneUniqueCheckPkRef struct {
	TimeWithTimeZoneUniqueCheckPkRef time.Time `db:"time_with_time_zone_unique_check_pk_ref"`
}
