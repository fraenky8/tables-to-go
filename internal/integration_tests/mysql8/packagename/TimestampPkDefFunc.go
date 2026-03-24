package models

import (
	"time"
)

type TimestampPkDefFunc struct {
	TimestampPkDefFunc time.Time `db:"timestamp_pk_def_func"`
}
