package dto

import (
	"time"
)

type DatePk struct {
	DatePk time.Time `db:"date_pk"`
}
