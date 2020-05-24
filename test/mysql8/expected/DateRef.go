package dto

import (
	"github.com/go-sql-driver/mysql"
)

type DateRef struct {
	DateRef mysql.NullTime `db:"date_ref"`
}
