package dto

import (
	"time"
)

type TimeWithTimeZonePk struct {
	TimeWithTimeZonePk time.Time `db:"time_with_time_zone_pk"`
}
