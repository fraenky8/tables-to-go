package dto

import (
	"time"
)

type DatetimeUniqueCheckPkRef struct {
	DatetimeUniqueCheckPkRef time.Time `db:"datetime_unique_check_pk_ref"`
}
