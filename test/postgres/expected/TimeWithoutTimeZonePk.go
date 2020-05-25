package dto

import (
	"time"
)

type TimeWithoutTimeZonePk struct {
	TimeWithoutTimeZonePk time.Time `db:"time_without_time_zone_pk"`
}
