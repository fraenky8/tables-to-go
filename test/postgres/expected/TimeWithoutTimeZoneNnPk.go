package dto

import (
	"time"
)

type TimeWithoutTimeZoneNnPk struct {
	TimeWithoutTimeZoneNnPk time.Time `db:"time_without_time_zone_nn_pk"`
}
