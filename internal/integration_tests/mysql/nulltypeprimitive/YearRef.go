package dto

import (
	"time"
)

type YearRef struct {
	YearRef *time.Time `db:"year_ref"`
}
