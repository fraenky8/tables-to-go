package dto

import (
	"time"
)

type TimestampWithoutTimeZonePkDefFunc struct {
	TimestampWithoutTimeZonePkDefFunc time.Time `db:"timestamp_without_time_zone_pk_def_func"`
}
