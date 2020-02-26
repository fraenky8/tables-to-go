package dto

import (
	"time"
)

type TimestamptzPkRef struct {
	TimestamptzPkRef time.Time `db:"timestamptz_pk_ref"`
}
