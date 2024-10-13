package dto

import (
	"time"
)

type DateNnPk struct {
	DateNnPk time.Time `db:"date_nn_pk"`
}
