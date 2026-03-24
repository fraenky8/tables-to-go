package models

import (
	"time"
)

type TimePk struct {
	TimePk time.Time `db:"time_pk"`
}
