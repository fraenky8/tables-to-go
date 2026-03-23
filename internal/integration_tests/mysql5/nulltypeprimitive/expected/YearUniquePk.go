package dto

import (
	"time"
)

type YearUniquePk struct {
	YearUniquePk time.Time `db:"year_unique_pk"`
}
