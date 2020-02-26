package dto

import (
	"time"
)

type TimestamptzCheckPk struct {
	TimestamptzCheckPk time.Time `db:"timestamptz_check_pk"`
}
