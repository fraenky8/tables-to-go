package dto

import (
	"time"
)

type YearTable struct {
	Year                    *time.Time `db:"year"`
	YearNn                  time.Time  `db:"year_nn"`
	YearNnUnique            time.Time  `db:"year_nn_unique"`
	YearNnCheck             time.Time  `db:"year_nn_check"`
	YearNnRef               time.Time  `db:"year_nn_ref"`
	YearNnDefConst          time.Time  `db:"year_nn_def_const"`
	YearNnDefFunc           time.Time  `db:"year_nn_def_func"`
	YearNnUniqueCheck       time.Time  `db:"year_nn_unique_check"`
	YearUnique              *time.Time `db:"year_unique"`
	YearUniqueCheck         *time.Time `db:"year_unique_check"`
	YearUniqueRef           *time.Time `db:"year_unique_ref"`
	YearUniqueDefConst      *time.Time `db:"year_unique_def_const"`
	YearUniqueDefFunc       *time.Time `db:"year_unique_def_func"`
	YearCheck               *time.Time `db:"year_check"`
	YearCheckRef            *time.Time `db:"year_check_ref"`
	YearCheckDefConst       *time.Time `db:"year_check_def_const"`
	YearCheckDefFunc        *time.Time `db:"year_check_def_func"`
	YearRef                 *time.Time `db:"year_ref"`
	YearRefUniqueCheck      *time.Time `db:"year_ref_unique_check"`
	YearDefConst            *time.Time `db:"year_def_const"`
	YearDefConstUniqueCheck *time.Time `db:"year_def_const_unique_check"`
	YearDefFunc             *time.Time `db:"year_def_func"`
	YearDefFuncUniqueCheck  *time.Time `db:"year_def_func_unique_check"`
}
