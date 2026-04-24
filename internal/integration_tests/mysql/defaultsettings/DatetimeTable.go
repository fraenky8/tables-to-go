package dto

import (
	"database/sql"
	"time"
)

type DatetimeTable struct {
	Datetime                    sql.NullTime `db:"datetime"`
	DatetimeNn                  time.Time    `db:"datetime_nn"`
	DatetimeNnUnique            time.Time    `db:"datetime_nn_unique"`
	DatetimeNnCheck             time.Time    `db:"datetime_nn_check"`
	DatetimeNnRef               time.Time    `db:"datetime_nn_ref"`
	DatetimeNnDefConst          time.Time    `db:"datetime_nn_def_const"`
	DatetimeNnDefFunc           time.Time    `db:"datetime_nn_def_func"`
	DatetimeNnUniqueCheck       time.Time    `db:"datetime_nn_unique_check"`
	DatetimeUnique              sql.NullTime `db:"datetime_unique"`
	DatetimeUniqueCheck         sql.NullTime `db:"datetime_unique_check"`
	DatetimeUniqueRef           sql.NullTime `db:"datetime_unique_ref"`
	DatetimeUniqueDefConst      sql.NullTime `db:"datetime_unique_def_const"`
	DatetimeUniqueDefFunc       sql.NullTime `db:"datetime_unique_def_func"`
	DatetimeCheck               sql.NullTime `db:"datetime_check"`
	DatetimeCheckRef            sql.NullTime `db:"datetime_check_ref"`
	DatetimeCheckDefConst       sql.NullTime `db:"datetime_check_def_const"`
	DatetimeCheckDefFunc        sql.NullTime `db:"datetime_check_def_func"`
	DatetimeRef                 sql.NullTime `db:"datetime_ref"`
	DatetimeRefUniqueCheck      sql.NullTime `db:"datetime_ref_unique_check"`
	DatetimeDefConst            sql.NullTime `db:"datetime_def_const"`
	DatetimeDefConstUniqueCheck sql.NullTime `db:"datetime_def_const_unique_check"`
	DatetimeDefFunc             sql.NullTime `db:"datetime_def_func"`
	DatetimeDefFuncUniqueCheck  sql.NullTime `db:"datetime_def_func_unique_check"`
}
