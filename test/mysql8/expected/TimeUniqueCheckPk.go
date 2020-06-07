package dto

import (
	"time"
)

type TimeUniqueCheckPk struct {
	TimeUniqueCheckPk time.Time `db:"time_unique_check_pk"`
}
