package dto

import (
	"time"
)

type TimestampWithTimeZonePkDefFunc struct {
	TimestampWithTimeZonePkDefFunc time.Time `db:"timestamp_with_time_zone_pk_def_func"`
}
