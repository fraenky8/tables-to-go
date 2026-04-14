package dto

import (
	"time"
)

type TimestamptzRef struct {
	TimestamptzRef *time.Time `db:"timestamptz_ref"`
}
