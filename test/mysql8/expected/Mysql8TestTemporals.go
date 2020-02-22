package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Mysql8TestTemporals struct {
	T    mysql.NullTime `db:"t"`
	TNn  time.Time      `db:"t_nn"`
	D    mysql.NullTime `db:"d"`
	DNn  time.Time      `db:"d_nn"`
	Dt   mysql.NullTime `db:"dt"`
	DtNn time.Time      `db:"dt_nn"`
	Ts   mysql.NullTime `db:"ts"`
	TsNn time.Time      `db:"ts_nn"`
}
