package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type DateTable struct {
	Date                    mysql.NullTime `db:"date"`
	DateNn                  time.Time      `db:"date_nn"`
	DateNnUnique            time.Time      `db:"date_nn_unique"`
	DateNnCheck             time.Time      `db:"date_nn_check"`
	DateNnRef               time.Time      `db:"date_nn_ref"`
	DateNnDefConst          time.Time      `db:"date_nn_def_const"`
	DateNnDefFunc           time.Time      `db:"date_nn_def_func"`
	DateNnUniqueCheck       time.Time      `db:"date_nn_unique_check"`
	DateUnique              mysql.NullTime `db:"date_unique"`
	DateUniqueCheck         mysql.NullTime `db:"date_unique_check"`
	DateUniqueRef           mysql.NullTime `db:"date_unique_ref"`
	DateUniqueDefConst      mysql.NullTime `db:"date_unique_def_const"`
	DateUniqueDefFunc       mysql.NullTime `db:"date_unique_def_func"`
	DateCheck               mysql.NullTime `db:"date_check"`
	DateCheckRef            mysql.NullTime `db:"date_check_ref"`
	DateCheckDefConst       mysql.NullTime `db:"date_check_def_const"`
	DateCheckDefFunc        mysql.NullTime `db:"date_check_def_func"`
	DateRef                 mysql.NullTime `db:"date_ref"`
	DateRefUniqueCheck      mysql.NullTime `db:"date_ref_unique_check"`
	DateDefConst            mysql.NullTime `db:"date_def_const"`
	DateDefConstUniqueCheck mysql.NullTime `db:"date_def_const_unique_check"`
	DateDefFunc             mysql.NullTime `db:"date_def_func"`
	DateDefFuncUniqueCheck  mysql.NullTime `db:"date_def_func_unique_check"`
}
