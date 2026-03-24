package dto

import (
	"time"
)

type TimestamptzPk struct {
	TimestamptzPk time.Time `db:"timestamptz_pk"`
}
