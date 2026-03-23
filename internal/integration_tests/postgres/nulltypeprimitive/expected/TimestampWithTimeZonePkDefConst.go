package dto

import (
	"time"
)

type TimestampWithTimeZonePkDefConst struct {
	TimestampWithTimeZonePkDefConst time.Time `db:"timestamp_with_time_zone_pk_def_const"`
}
