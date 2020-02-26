package dto

import (
	"time"
)

type TimestampWithTimeZoneNnPk struct {
	TimestampWithTimeZoneNnPk time.Time `db:"timestamp_with_time_zone_nn_pk"`
}
