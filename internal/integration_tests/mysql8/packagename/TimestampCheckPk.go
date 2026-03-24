package models

import (
	"time"
)

type TimestampCheckPk struct {
	TimestampCheckPk time.Time `db:"timestamp_check_pk"`
}
