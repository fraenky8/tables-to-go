package dto

import (
	"time"
)

type TimestampUniqueCheckPk struct {
	TimestampUniqueCheckPk time.Time `db:"timestamp_unique_check_pk"`
}
