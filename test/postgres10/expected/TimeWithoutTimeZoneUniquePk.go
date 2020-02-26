package dto

import (
	"time"
)

type TimeWithoutTimeZoneUniquePk struct {
	TimeWithoutTimeZoneUniquePk time.Time `db:"time_without_time_zone_unique_pk"`
}
