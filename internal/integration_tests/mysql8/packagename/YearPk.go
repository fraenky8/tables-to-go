package models

import (
	"time"
)

type YearPk struct {
	YearPk time.Time `db:"year_pk"`
}
