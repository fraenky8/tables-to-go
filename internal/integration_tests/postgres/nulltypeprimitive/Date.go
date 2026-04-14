package dto

import (
	"time"
)

type Date struct {
	Date                    *time.Time `db:"date"`
	DateNn                  time.Time  `db:"date_nn"`
	DateNnUnique            time.Time  `db:"date_nn_unique"`
	DateNnCheck             time.Time  `db:"date_nn_check"`
	DateNnRef               time.Time  `db:"date_nn_ref"`
	DateNnDefConst          time.Time  `db:"date_nn_def_const"`
	DateNnDefFunc           time.Time  `db:"date_nn_def_func"`
	DateNnUniqueCheck       time.Time  `db:"date_nn_unique_check"`
	DateUnique              *time.Time `db:"date_unique"`
	DateUniqueCheck         *time.Time `db:"date_unique_check"`
	DateUniqueRef           *time.Time `db:"date_unique_ref"`
	DateUniqueDefConst      *time.Time `db:"date_unique_def_const"`
	DateUniqueDefFunc       *time.Time `db:"date_unique_def_func"`
	DateCheck               *time.Time `db:"date_check"`
	DateCheckRef            *time.Time `db:"date_check_ref"`
	DateCheckDefConst       *time.Time `db:"date_check_def_const"`
	DateCheckDefFunc        *time.Time `db:"date_check_def_func"`
	DateRef                 *time.Time `db:"date_ref"`
	DateRefDefConst         *time.Time `db:"date_ref_def_const"`
	DateRefDefFunc          *time.Time `db:"date_ref_def_func"`
	DateRefUniqueCheck      *time.Time `db:"date_ref_unique_check"`
	DateDefConst            *time.Time `db:"date_def_const"`
	DateDefConstUniqueCheck *time.Time `db:"date_def_const_unique_check"`
	DateDefFunc             *time.Time `db:"date_def_func"`
	DateDefFuncUniqueCheck  *time.Time `db:"date_def_func_unique_check"`
}
