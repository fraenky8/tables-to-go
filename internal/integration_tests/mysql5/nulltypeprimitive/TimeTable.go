package dto

import (
	"time"
)

type TimeTable struct {
	Time                    *time.Time `db:"time"`
	TimeNn                  time.Time  `db:"time_nn"`
	TimeNnUnique            time.Time  `db:"time_nn_unique"`
	TimeNnCheck             time.Time  `db:"time_nn_check"`
	TimeNnRef               time.Time  `db:"time_nn_ref"`
	TimeNnDefConst          time.Time  `db:"time_nn_def_const"`
	TimeNnDefFunc           time.Time  `db:"time_nn_def_func"`
	TimeNnUniqueCheck       time.Time  `db:"time_nn_unique_check"`
	TimeUnique              *time.Time `db:"time_unique"`
	TimeUniqueCheck         *time.Time `db:"time_unique_check"`
	TimeUniqueRef           *time.Time `db:"time_unique_ref"`
	TimeUniqueDefConst      *time.Time `db:"time_unique_def_const"`
	TimeUniqueDefFunc       *time.Time `db:"time_unique_def_func"`
	TimeCheck               *time.Time `db:"time_check"`
	TimeCheckRef            *time.Time `db:"time_check_ref"`
	TimeCheckDefConst       *time.Time `db:"time_check_def_const"`
	TimeCheckDefFunc        *time.Time `db:"time_check_def_func"`
	TimeRef                 *time.Time `db:"time_ref"`
	TimeRefUniqueCheck      *time.Time `db:"time_ref_unique_check"`
	TimeDefConst            *time.Time `db:"time_def_const"`
	TimeDefConstUniqueCheck *time.Time `db:"time_def_const_unique_check"`
	TimeDefFunc             *time.Time `db:"time_def_func"`
	TimeDefFuncUniqueCheck  *time.Time `db:"time_def_func_unique_check"`
}
