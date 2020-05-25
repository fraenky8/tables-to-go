package dto

import (
	"time"
)

type TimestamptzUniqueCheckPk struct {
	TimestamptzUniqueCheckPk time.Time `db:"timestamptz_unique_check_pk"`
}
