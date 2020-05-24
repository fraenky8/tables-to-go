package dto

import (
	"github.com/go-sql-driver/mysql"
)

type TimeRef struct {
	TimeRef mysql.NullTime `db:"time_ref"`
}
