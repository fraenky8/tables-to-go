package dto

import (
	"time"
)

type TimestampUniqueCheckPkRef struct {
	TimestampUniqueCheckPkRef time.Time `db:"timestamp_unique_check_pk_ref"`
}
