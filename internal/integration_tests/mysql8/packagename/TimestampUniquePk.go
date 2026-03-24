package models

import (
	"time"
)

type TimestampUniquePk struct {
	TimestampUniquePk time.Time `db:"timestamp_unique_pk"`
}
