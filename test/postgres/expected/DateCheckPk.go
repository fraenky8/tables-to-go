package dto

import (
	"time"
)

type DateCheckPk struct {
	DateCheckPk time.Time `db:"date_check_pk"`
}
