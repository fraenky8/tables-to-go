package dto

import (
	"time"
)

type DatetimePk struct {
	DatetimePk time.Time `db:"datetime_pk"`
}
