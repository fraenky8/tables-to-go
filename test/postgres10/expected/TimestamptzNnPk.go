package dto

import (
	"time"
)

type TimestamptzNnPk struct {
	TimestamptzNnPk time.Time `db:"timestamptz_nn_pk"`
}
