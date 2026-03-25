package dto

import (
	"time"
)

type TimeWithoutTimeZoneRef struct {
	TimeWithoutTimeZoneRef *time.Time `db:"time_without_time_zone_ref"`
}
