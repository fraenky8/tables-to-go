package models

import (
	"time"
)

type DatetimeNnPk struct {
	DatetimeNnPk time.Time `db:"datetime_nn_pk"`
}
