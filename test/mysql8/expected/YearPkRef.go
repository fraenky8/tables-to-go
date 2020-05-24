package dto

import (
	"time"
)

type YearPkRef struct {
	YearPkRef time.Time `db:"year_pk_ref"`
}
