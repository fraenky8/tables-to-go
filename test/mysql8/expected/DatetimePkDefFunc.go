package dto

import (
	"time"
)

type DatetimePkDefFunc struct {
	DatetimePkDefFunc time.Time `db:"datetime_pk_def_func"`
}
