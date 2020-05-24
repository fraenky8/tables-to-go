package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type DatetimeTable struct {
	Datetime                    mysql.NullTime `db:"datetime"`
	DatetimeNn                  time.Time      `db:"datetime_nn"`
	DatetimeNnUnique            time.Time      `db:"datetime_nn_unique"`
	DatetimeNnCheck             time.Time      `db:"datetime_nn_check"`
	DatetimeNnRef               time.Time      `db:"datetime_nn_ref"`
	DatetimeNnDefConst          time.Time      `db:"datetime_nn_def_const"`
	DatetimeNnDefFunc           time.Time      `db:"datetime_nn_def_func"`
	DatetimeNnUniqueCheck       time.Time      `db:"datetime_nn_unique_check"`
	DatetimeUnique              mysql.NullTime `db:"datetime_unique"`
	DatetimeUniqueCheck         mysql.NullTime `db:"datetime_unique_check"`
	DatetimeUniqueRef           mysql.NullTime `db:"datetime_unique_ref"`
	DatetimeUniqueDefConst      mysql.NullTime `db:"datetime_unique_def_const"`
	DatetimeUniqueDefFunc       mysql.NullTime `db:"datetime_unique_def_func"`
	DatetimeCheck               mysql.NullTime `db:"datetime_check"`
	DatetimeCheckRef            mysql.NullTime `db:"datetime_check_ref"`
	DatetimeCheckDefConst       mysql.NullTime `db:"datetime_check_def_const"`
	DatetimeCheckDefFunc        mysql.NullTime `db:"datetime_check_def_func"`
	DatetimeRef                 mysql.NullTime `db:"datetime_ref"`
	DatetimeRefUniqueCheck      mysql.NullTime `db:"datetime_ref_unique_check"`
	DatetimeDefConst            mysql.NullTime `db:"datetime_def_const"`
	DatetimeDefConstUniqueCheck mysql.NullTime `db:"datetime_def_const_unique_check"`
	DatetimeDefFunc             mysql.NullTime `db:"datetime_def_func"`
	DatetimeDefFuncUniqueCheck  mysql.NullTime `db:"datetime_def_func_unique_check"`
}
