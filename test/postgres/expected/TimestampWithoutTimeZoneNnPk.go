package dto

import (
	"time"
)

type TimestampWithoutTimeZoneNnPk struct {
	TimestampWithoutTimeZoneNnPk time.Time `db:"timestamp_without_time_zone_nn_pk"`
}
