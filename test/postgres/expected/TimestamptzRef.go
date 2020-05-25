package dto

import (
	pg "github.com/lib/pq"
)

type TimestamptzRef struct {
	TimestamptzRef pg.NullTime `db:"timestamptz_ref"`
}
