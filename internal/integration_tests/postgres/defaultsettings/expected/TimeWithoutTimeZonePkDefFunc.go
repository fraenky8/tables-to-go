package dto

import (
	"time"
)

type TimeWithoutTimeZonePkDefFunc struct {
	TimeWithoutTimeZonePkDefFunc time.Time `db:"time_without_time_zone_pk_def_func"`
}
