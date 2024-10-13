package dto

import (
	"time"
)

type TimeWithoutTimeZonePkDefConst struct {
	TimeWithoutTimeZonePkDefConst time.Time `db:"time_without_time_zone_pk_def_const"`
}
