package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type TimeTable struct {
	Time                    mysql.NullTime `db:"time"`
	TimeNn                  time.Time      `db:"time_nn"`
	TimeNnUnique            time.Time      `db:"time_nn_unique"`
	TimeNnCheck             time.Time      `db:"time_nn_check"`
	TimeNnRef               time.Time      `db:"time_nn_ref"`
	TimeNnDefConst          time.Time      `db:"time_nn_def_const"`
	TimeNnDefFunc           time.Time      `db:"time_nn_def_func"`
	TimeNnUniqueCheck       time.Time      `db:"time_nn_unique_check"`
	TimeUnique              mysql.NullTime `db:"time_unique"`
	TimeUniqueCheck         mysql.NullTime `db:"time_unique_check"`
	TimeUniqueRef           mysql.NullTime `db:"time_unique_ref"`
	TimeUniqueDefConst      mysql.NullTime `db:"time_unique_def_const"`
	TimeUniqueDefFunc       mysql.NullTime `db:"time_unique_def_func"`
	TimeCheck               mysql.NullTime `db:"time_check"`
	TimeCheckRef            mysql.NullTime `db:"time_check_ref"`
	TimeCheckDefConst       mysql.NullTime `db:"time_check_def_const"`
	TimeCheckDefFunc        mysql.NullTime `db:"time_check_def_func"`
	TimeRef                 mysql.NullTime `db:"time_ref"`
	TimeRefUniqueCheck      mysql.NullTime `db:"time_ref_unique_check"`
	TimeDefConst            mysql.NullTime `db:"time_def_const"`
	TimeDefConstUniqueCheck mysql.NullTime `db:"time_def_const_unique_check"`
	TimeDefFunc             mysql.NullTime `db:"time_def_func"`
	TimeDefFuncUniqueCheck  mysql.NullTime `db:"time_def_func_unique_check"`
}
