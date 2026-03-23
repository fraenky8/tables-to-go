package dto

import (
	"time"
)

type TimestampPk struct {
	TimestampPk time.Time `db:"timestamp_pk"`
}
