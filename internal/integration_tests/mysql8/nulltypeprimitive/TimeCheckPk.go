package dto

import (
	"time"
)

type TimeCheckPk struct {
	TimeCheckPk time.Time `db:"time_check_pk"`
}
