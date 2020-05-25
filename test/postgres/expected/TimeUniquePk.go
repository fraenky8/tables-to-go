package dto

import (
	"time"
)

type TimeUniquePk struct {
	TimeUniquePk time.Time `db:"time_unique_pk"`
}
