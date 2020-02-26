package dto

import (
	pg "github.com/lib/pq"
)

type DateRef struct {
	DateRef pg.NullTime `db:"date_ref"`
}
