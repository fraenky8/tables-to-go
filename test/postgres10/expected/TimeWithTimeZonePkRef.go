package dto

import (
	"time"
)

type TimeWithTimeZonePkRef struct {
	TimeWithTimeZonePkRef time.Time `db:"time_with_time_zone_pk_ref"`
}
