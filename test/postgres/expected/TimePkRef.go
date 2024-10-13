package dto

import (
	"time"
)

type TimePkRef struct {
	TimePkRef time.Time `db:"time_pk_ref"`
}
