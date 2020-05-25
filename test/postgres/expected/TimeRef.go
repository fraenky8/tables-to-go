package dto

import (
	pg "github.com/lib/pq"
)

type TimeRef struct {
	TimeRef pg.NullTime `db:"time_ref"`
}
