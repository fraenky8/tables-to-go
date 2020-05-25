package dto

import (
	"time"
)

type TimestamptzPkDefFunc struct {
	TimestamptzPkDefFunc time.Time `db:"timestamptz_pk_def_func"`
}
