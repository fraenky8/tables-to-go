package dto

import (
	"time"
)

type DatetimeUniquePk struct {
	DatetimeUniquePk time.Time `db:"datetime_unique_pk"`
}
