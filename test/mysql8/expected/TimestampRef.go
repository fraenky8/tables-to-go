package dto

import (
	"github.com/go-sql-driver/mysql"
)

type TimestampRef struct {
	TimestampRef mysql.NullTime `db:"timestamp_ref"`
}
