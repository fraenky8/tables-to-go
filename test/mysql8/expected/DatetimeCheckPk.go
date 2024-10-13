package dto

import (
	"time"
)

type DatetimeCheckPk struct {
	DatetimeCheckPk time.Time `db:"datetime_check_pk"`
}
