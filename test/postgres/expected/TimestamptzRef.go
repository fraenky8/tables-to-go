package dto

import (
	"database/sql"
)

type TimestamptzRef struct {
	TimestamptzRef sql.NullTime `db:"timestamptz_ref"`
}
