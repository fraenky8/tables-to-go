package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type Timestamp struct {
	Timestamp                    pg.NullTime `db:"timestamp"`
	TimestampNn                  time.Time   `db:"timestamp_nn"`
	TimestampNnUnique            time.Time   `db:"timestamp_nn_unique"`
	TimestampNnCheck             time.Time   `db:"timestamp_nn_check"`
	TimestampNnRef               time.Time   `db:"timestamp_nn_ref"`
	TimestampNnDefConst          time.Time   `db:"timestamp_nn_def_const"`
	TimestampNnDefFunc           time.Time   `db:"timestamp_nn_def_func"`
	TimestampNnUniqueCheck       time.Time   `db:"timestamp_nn_unique_check"`
	TimestampUnique              pg.NullTime `db:"timestamp_unique"`
	TimestampUniqueCheck         pg.NullTime `db:"timestamp_unique_check"`
	TimestampUniqueRef           pg.NullTime `db:"timestamp_unique_ref"`
	TimestampUniqueDefConst      pg.NullTime `db:"timestamp_unique_def_const"`
	TimestampUniqueDefFunc       pg.NullTime `db:"timestamp_unique_def_func"`
	TimestampCheck               pg.NullTime `db:"timestamp_check"`
	TimestampCheckRef            pg.NullTime `db:"timestamp_check_ref"`
	TimestampCheckDefConst       pg.NullTime `db:"timestamp_check_def_const"`
	TimestampCheckDefFunc        pg.NullTime `db:"timestamp_check_def_func"`
	TimestampRef                 pg.NullTime `db:"timestamp_ref"`
	TimestampRefDefConst         pg.NullTime `db:"timestamp_ref_def_const"`
	TimestampRefDefFunc          pg.NullTime `db:"timestamp_ref_def_func"`
	TimestampRefUniqueCheck      pg.NullTime `db:"timestamp_ref_unique_check"`
	TimestampDefConst            pg.NullTime `db:"timestamp_def_const"`
	TimestampDefConstUniqueCheck pg.NullTime `db:"timestamp_def_const_unique_check"`
	TimestampDefFunc             pg.NullTime `db:"timestamp_def_func"`
	TimestampDefFuncUniqueCheck  pg.NullTime `db:"timestamp_def_func_unique_check"`
}
