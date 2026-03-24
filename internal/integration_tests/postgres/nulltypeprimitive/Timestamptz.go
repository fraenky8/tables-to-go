package dto

import (
	"time"
)

type Timestamptz struct {
	Timestamptz                    *time.Time `db:"timestamptz"`
	TimestamptzNn                  time.Time  `db:"timestamptz_nn"`
	TimestamptzNnUnique            time.Time  `db:"timestamptz_nn_unique"`
	TimestamptzNnCheck             time.Time  `db:"timestamptz_nn_check"`
	TimestamptzNnRef               time.Time  `db:"timestamptz_nn_ref"`
	TimestamptzNnDefConst          time.Time  `db:"timestamptz_nn_def_const"`
	TimestamptzNnDefFunc           time.Time  `db:"timestamptz_nn_def_func"`
	TimestamptzNnUniqueCheck       time.Time  `db:"timestamptz_nn_unique_check"`
	TimestamptzUnique              *time.Time `db:"timestamptz_unique"`
	TimestamptzUniqueCheck         *time.Time `db:"timestamptz_unique_check"`
	TimestamptzUniqueRef           *time.Time `db:"timestamptz_unique_ref"`
	TimestamptzUniqueDefConst      *time.Time `db:"timestamptz_unique_def_const"`
	TimestamptzUniqueDefFunc       *time.Time `db:"timestamptz_unique_def_func"`
	TimestamptzCheck               *time.Time `db:"timestamptz_check"`
	TimestamptzCheckRef            *time.Time `db:"timestamptz_check_ref"`
	TimestamptzCheckDefConst       *time.Time `db:"timestamptz_check_def_const"`
	TimestamptzCheckDefFunc        *time.Time `db:"timestamptz_check_def_func"`
	TimestamptzRef                 *time.Time `db:"timestamptz_ref"`
	TimestamptzRefDefConst         *time.Time `db:"timestamptz_ref_def_const"`
	TimestamptzRefDefFunc          *time.Time `db:"timestamptz_ref_def_func"`
	TimestamptzRefUniqueCheck      *time.Time `db:"timestamptz_ref_unique_check"`
	TimestamptzDefConst            *time.Time `db:"timestamptz_def_const"`
	TimestamptzDefConstUniqueCheck *time.Time `db:"timestamptz_def_const_unique_check"`
	TimestamptzDefFunc             *time.Time `db:"timestamptz_def_func"`
	TimestamptzDefFuncUniqueCheck  *time.Time `db:"timestamptz_def_func_unique_check"`
}
