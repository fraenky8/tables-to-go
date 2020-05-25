package dto

import (
	"time"
)

type TimestamptzUniqueCheckPkRef struct {
	TimestamptzUniqueCheckPkRef time.Time `db:"timestamptz_unique_check_pk_ref"`
}
