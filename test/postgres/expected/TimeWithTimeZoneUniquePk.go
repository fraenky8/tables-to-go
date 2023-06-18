package dto

import (
	"time"
)

type TimeWithTimeZoneUniquePk struct {
	TimeWithTimeZoneUniquePk time.Time `db:"time_with_time_zone_unique_pk"`
}
