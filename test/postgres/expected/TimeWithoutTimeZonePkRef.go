package dto

import (
	"time"
)

type TimeWithoutTimeZonePkRef struct {
	TimeWithoutTimeZonePkRef time.Time `db:"time_without_time_zone_pk_ref"`
}
