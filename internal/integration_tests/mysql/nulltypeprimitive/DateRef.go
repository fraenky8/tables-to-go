package dto

import (
	"time"
)

type DateRef struct {
	DateRef *time.Time `db:"date_ref"`
}
