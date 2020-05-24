package dto

import (
	"github.com/go-sql-driver/mysql"
)

type YearRef struct {
	YearRef mysql.NullTime `db:"year_ref"`
}
