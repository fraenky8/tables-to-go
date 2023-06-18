package dto

import (
	"time"
)

type TimeWithoutTimeZoneNnUniqueCheckPk struct {
	TimeWithoutTimeZoneNnUniqueCheckPk time.Time `db:"time_without_time_zone_nn_unique_check_pk"`
}
