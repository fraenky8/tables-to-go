package dto

import (
	"time"
)

type DatetimeRef struct {
	DatetimeRef *time.Time `db:"datetime_ref"`
}
