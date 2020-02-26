package dto

import (
	"time"
)

type TimeWithTimeZonePkDefFunc struct {
	TimeWithTimeZonePkDefFunc time.Time `db:"time_with_time_zone_pk_def_func"`
}
