package models

import (
	"time"
)

type TimestampNnPk struct {
	TimestampNnPk time.Time `db:"timestamp_nn_pk"`
}
