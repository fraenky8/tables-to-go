package dto

import (
	"time"
)

type TimeUniqueCheckPkRef struct {
	TimeUniqueCheckPkRef time.Time `db:"time_unique_check_pk_ref"`
}
