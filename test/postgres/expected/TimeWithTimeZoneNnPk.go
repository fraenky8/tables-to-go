package dto

import (
	"time"
)

type TimeWithTimeZoneNnPk struct {
	TimeWithTimeZoneNnPk time.Time `db:"time_with_time_zone_nn_pk"`
}
