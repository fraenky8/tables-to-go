package dto

import (
	"time"
)

type DateUniquePk struct {
	DateUniquePk time.Time `db:"date_unique_pk"`
}
