package dto

import (
	"time"
)

type DatePkRef struct {
	DatePkRef time.Time `db:"date_pk_ref"`
}
