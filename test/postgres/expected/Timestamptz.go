package dto

import (
	"database/sql"
	"time"
)

type Timestamptz struct {
	Timestamptz                    sql.NullTime `db:"timestamptz"`
	TimestamptzNn                  time.Time    `db:"timestamptz_nn"`
	TimestamptzNnUnique            time.Time    `db:"timestamptz_nn_unique"`
	TimestamptzNnCheck             time.Time    `db:"timestamptz_nn_check"`
	TimestamptzNnRef               time.Time    `db:"timestamptz_nn_ref"`
	TimestamptzNnDefConst          time.Time    `db:"timestamptz_nn_def_const"`
	TimestamptzNnDefFunc           time.Time    `db:"timestamptz_nn_def_func"`
	TimestamptzNnUniqueCheck       time.Time    `db:"timestamptz_nn_unique_check"`
	TimestamptzUnique              sql.NullTime `db:"timestamptz_unique"`
	TimestamptzUniqueCheck         sql.NullTime `db:"timestamptz_unique_check"`
	TimestamptzUniqueRef           sql.NullTime `db:"timestamptz_unique_ref"`
	TimestamptzUniqueDefConst      sql.NullTime `db:"timestamptz_unique_def_const"`
	TimestamptzUniqueDefFunc       sql.NullTime `db:"timestamptz_unique_def_func"`
	TimestamptzCheck               sql.NullTime `db:"timestamptz_check"`
	TimestamptzCheckRef            sql.NullTime `db:"timestamptz_check_ref"`
	TimestamptzCheckDefConst       sql.NullTime `db:"timestamptz_check_def_const"`
	TimestamptzCheckDefFunc        sql.NullTime `db:"timestamptz_check_def_func"`
	TimestamptzRef                 sql.NullTime `db:"timestamptz_ref"`
	TimestamptzRefDefConst         sql.NullTime `db:"timestamptz_ref_def_const"`
	TimestamptzRefDefFunc          sql.NullTime `db:"timestamptz_ref_def_func"`
	TimestamptzRefUniqueCheck      sql.NullTime `db:"timestamptz_ref_unique_check"`
	TimestamptzDefConst            sql.NullTime `db:"timestamptz_def_const"`
	TimestamptzDefConstUniqueCheck sql.NullTime `db:"timestamptz_def_const_unique_check"`
	TimestamptzDefFunc             sql.NullTime `db:"timestamptz_def_func"`
	TimestamptzDefFuncUniqueCheck  sql.NullTime `db:"timestamptz_def_func_unique_check"`
}
