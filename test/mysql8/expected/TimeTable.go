package dto

import (
	"database/sql"
	"time"
)

type TimeTable struct {
	Time                    sql.NullTime `db:"time"`
	TimeNn                  time.Time    `db:"time_nn"`
	TimeNnUnique            time.Time    `db:"time_nn_unique"`
	TimeNnCheck             time.Time    `db:"time_nn_check"`
	TimeNnRef               time.Time    `db:"time_nn_ref"`
	TimeNnDefConst          time.Time    `db:"time_nn_def_const"`
	TimeNnDefFunc           time.Time    `db:"time_nn_def_func"`
	TimeNnUniqueCheck       time.Time    `db:"time_nn_unique_check"`
	TimeUnique              sql.NullTime `db:"time_unique"`
	TimeUniqueCheck         sql.NullTime `db:"time_unique_check"`
	TimeUniqueRef           sql.NullTime `db:"time_unique_ref"`
	TimeUniqueDefConst      sql.NullTime `db:"time_unique_def_const"`
	TimeUniqueDefFunc       sql.NullTime `db:"time_unique_def_func"`
	TimeCheck               sql.NullTime `db:"time_check"`
	TimeCheckRef            sql.NullTime `db:"time_check_ref"`
	TimeCheckDefConst       sql.NullTime `db:"time_check_def_const"`
	TimeCheckDefFunc        sql.NullTime `db:"time_check_def_func"`
	TimeRef                 sql.NullTime `db:"time_ref"`
	TimeRefUniqueCheck      sql.NullTime `db:"time_ref_unique_check"`
	TimeDefConst            sql.NullTime `db:"time_def_const"`
	TimeDefConstUniqueCheck sql.NullTime `db:"time_def_const_unique_check"`
	TimeDefFunc             sql.NullTime `db:"time_def_func"`
	TimeDefFuncUniqueCheck  sql.NullTime `db:"time_def_func_unique_check"`
}
