package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type Timestamptz struct {
	Timestamptz                    pg.NullTime `db:"timestamptz"`
	TimestamptzNn                  time.Time   `db:"timestamptz_nn"`
	TimestamptzNnUnique            time.Time   `db:"timestamptz_nn_unique"`
	TimestamptzNnCheck             time.Time   `db:"timestamptz_nn_check"`
	TimestamptzNnRef               time.Time   `db:"timestamptz_nn_ref"`
	TimestamptzNnDefConst          time.Time   `db:"timestamptz_nn_def_const"`
	TimestamptzNnDefFunc           time.Time   `db:"timestamptz_nn_def_func"`
	TimestamptzNnUniqueCheck       time.Time   `db:"timestamptz_nn_unique_check"`
	TimestamptzUnique              pg.NullTime `db:"timestamptz_unique"`
	TimestamptzUniqueCheck         pg.NullTime `db:"timestamptz_unique_check"`
	TimestamptzUniqueRef           pg.NullTime `db:"timestamptz_unique_ref"`
	TimestamptzUniqueDefConst      pg.NullTime `db:"timestamptz_unique_def_const"`
	TimestamptzUniqueDefFunc       pg.NullTime `db:"timestamptz_unique_def_func"`
	TimestamptzCheck               pg.NullTime `db:"timestamptz_check"`
	TimestamptzCheckRef            pg.NullTime `db:"timestamptz_check_ref"`
	TimestamptzCheckDefConst       pg.NullTime `db:"timestamptz_check_def_const"`
	TimestamptzCheckDefFunc        pg.NullTime `db:"timestamptz_check_def_func"`
	TimestamptzRef                 pg.NullTime `db:"timestamptz_ref"`
	TimestamptzRefDefConst         pg.NullTime `db:"timestamptz_ref_def_const"`
	TimestamptzRefDefFunc          pg.NullTime `db:"timestamptz_ref_def_func"`
	TimestamptzRefUniqueCheck      pg.NullTime `db:"timestamptz_ref_unique_check"`
	TimestamptzDefConst            pg.NullTime `db:"timestamptz_def_const"`
	TimestamptzDefConstUniqueCheck pg.NullTime `db:"timestamptz_def_const_unique_check"`
	TimestamptzDefFunc             pg.NullTime `db:"timestamptz_def_func"`
	TimestamptzDefFuncUniqueCheck  pg.NullTime `db:"timestamptz_def_func_unique_check"`
}
