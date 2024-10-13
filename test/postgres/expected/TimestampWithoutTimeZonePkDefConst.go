package dto

import (
	"time"
)

type TimestampWithoutTimeZonePkDefConst struct {
	TimestampWithoutTimeZonePkDefConst time.Time `db:"timestamp_without_time_zone_pk_def_const"`
}
