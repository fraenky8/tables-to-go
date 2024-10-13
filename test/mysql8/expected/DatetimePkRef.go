package dto

import (
	"time"
)

type DatetimePkRef struct {
	DatetimePkRef time.Time `db:"datetime_pk_ref"`
}
