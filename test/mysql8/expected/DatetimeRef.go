package dto

import (
	"github.com/go-sql-driver/mysql"
)

type DatetimeRef struct {
	DatetimeRef mysql.NullTime `db:"datetime_ref"`
}
