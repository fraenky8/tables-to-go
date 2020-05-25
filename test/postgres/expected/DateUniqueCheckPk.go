package dto

import (
	"time"
)

type DateUniqueCheckPk struct {
	DateUniqueCheckPk time.Time `db:"date_unique_check_pk"`
}
