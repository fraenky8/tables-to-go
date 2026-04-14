package dto

import (
	"time"
)

type TimeRef struct {
	TimeRef *time.Time `db:"time_ref"`
}
