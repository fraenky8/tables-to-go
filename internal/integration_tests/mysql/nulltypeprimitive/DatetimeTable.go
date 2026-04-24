package dto

import (
	"time"
)

type DatetimeTable struct {
	Datetime                    *time.Time `db:"datetime"`
	DatetimeNn                  time.Time  `db:"datetime_nn"`
	DatetimeNnUnique            time.Time  `db:"datetime_nn_unique"`
	DatetimeNnCheck             time.Time  `db:"datetime_nn_check"`
	DatetimeNnRef               time.Time  `db:"datetime_nn_ref"`
	DatetimeNnDefConst          time.Time  `db:"datetime_nn_def_const"`
	DatetimeNnDefFunc           time.Time  `db:"datetime_nn_def_func"`
	DatetimeNnUniqueCheck       time.Time  `db:"datetime_nn_unique_check"`
	DatetimeUnique              *time.Time `db:"datetime_unique"`
	DatetimeUniqueCheck         *time.Time `db:"datetime_unique_check"`
	DatetimeUniqueRef           *time.Time `db:"datetime_unique_ref"`
	DatetimeUniqueDefConst      *time.Time `db:"datetime_unique_def_const"`
	DatetimeUniqueDefFunc       *time.Time `db:"datetime_unique_def_func"`
	DatetimeCheck               *time.Time `db:"datetime_check"`
	DatetimeCheckRef            *time.Time `db:"datetime_check_ref"`
	DatetimeCheckDefConst       *time.Time `db:"datetime_check_def_const"`
	DatetimeCheckDefFunc        *time.Time `db:"datetime_check_def_func"`
	DatetimeRef                 *time.Time `db:"datetime_ref"`
	DatetimeRefUniqueCheck      *time.Time `db:"datetime_ref_unique_check"`
	DatetimeDefConst            *time.Time `db:"datetime_def_const"`
	DatetimeDefConstUniqueCheck *time.Time `db:"datetime_def_const_unique_check"`
	DatetimeDefFunc             *time.Time `db:"datetime_def_func"`
	DatetimeDefFuncUniqueCheck  *time.Time `db:"datetime_def_func_unique_check"`
}
