package dto

import (
	"time"
)

type TimeNnPk struct {
	TimeNnPk time.Time `db:"time_nn_pk"`
}
