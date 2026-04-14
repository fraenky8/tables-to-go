package dto

import (
	"time"
)

type TimestamptzUniquePk struct {
	TimestamptzUniquePk time.Time `db:"timestamptz_unique_pk"`
}
