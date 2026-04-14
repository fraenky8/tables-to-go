package dto

import (
	"time"
)

type TimestampTable struct {
	Timestamp                    *time.Time `db:"timestamp"`
	TimestampNn                  time.Time  `db:"timestamp_nn"`
	TimestampNnUnique            time.Time  `db:"timestamp_nn_unique"`
	TimestampNnCheck             time.Time  `db:"timestamp_nn_check"`
	TimestampNnRef               time.Time  `db:"timestamp_nn_ref"`
	TimestampNnDefConst          time.Time  `db:"timestamp_nn_def_const"`
	TimestampNnDefFunc           time.Time  `db:"timestamp_nn_def_func"`
	TimestampNnUniqueCheck       time.Time  `db:"timestamp_nn_unique_check"`
	TimestampUnique              *time.Time `db:"timestamp_unique"`
	TimestampUniqueCheck         *time.Time `db:"timestamp_unique_check"`
	TimestampUniqueRef           *time.Time `db:"timestamp_unique_ref"`
	TimestampUniqueDefConst      *time.Time `db:"timestamp_unique_def_const"`
	TimestampUniqueDefFunc       *time.Time `db:"timestamp_unique_def_func"`
	TimestampCheck               *time.Time `db:"timestamp_check"`
	TimestampCheckRef            *time.Time `db:"timestamp_check_ref"`
	TimestampCheckDefConst       *time.Time `db:"timestamp_check_def_const"`
	TimestampCheckDefFunc        *time.Time `db:"timestamp_check_def_func"`
	TimestampRef                 *time.Time `db:"timestamp_ref"`
	TimestampRefUniqueCheck      *time.Time `db:"timestamp_ref_unique_check"`
	TimestampDefConst            *time.Time `db:"timestamp_def_const"`
	TimestampDefConstUniqueCheck *time.Time `db:"timestamp_def_const_unique_check"`
	TimestampDefFunc             *time.Time `db:"timestamp_def_func"`
	TimestampDefFuncUniqueCheck  *time.Time `db:"timestamp_def_func_unique_check"`
}
