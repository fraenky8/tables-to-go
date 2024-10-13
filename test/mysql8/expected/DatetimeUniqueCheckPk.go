package dto

import (
	"time"
)

type DatetimeUniqueCheckPk struct {
	DatetimeUniqueCheckPk time.Time `db:"datetime_unique_check_pk"`
}
