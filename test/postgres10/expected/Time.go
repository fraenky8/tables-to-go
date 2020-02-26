package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type Time struct {
	Time                    pg.NullTime `db:"time"`
	TimeNn                  time.Time   `db:"time_nn"`
	TimeNnUnique            time.Time   `db:"time_nn_unique"`
	TimeNnCheck             time.Time   `db:"time_nn_check"`
	TimeNnRef               time.Time   `db:"time_nn_ref"`
	TimeNnDefConst          time.Time   `db:"time_nn_def_const"`
	TimeNnDefFunc           time.Time   `db:"time_nn_def_func"`
	TimeNnUniqueCheck       time.Time   `db:"time_nn_unique_check"`
	TimeUnique              pg.NullTime `db:"time_unique"`
	TimeUniqueCheck         pg.NullTime `db:"time_unique_check"`
	TimeUniqueRef           pg.NullTime `db:"time_unique_ref"`
	TimeUniqueDefConst      pg.NullTime `db:"time_unique_def_const"`
	TimeUniqueDefFunc       pg.NullTime `db:"time_unique_def_func"`
	TimeCheck               pg.NullTime `db:"time_check"`
	TimeCheckRef            pg.NullTime `db:"time_check_ref"`
	TimeCheckDefConst       pg.NullTime `db:"time_check_def_const"`
	TimeCheckDefFunc        pg.NullTime `db:"time_check_def_func"`
	TimeRef                 pg.NullTime `db:"time_ref"`
	TimeRefDefConst         pg.NullTime `db:"time_ref_def_const"`
	TimeRefDefFunc          pg.NullTime `db:"time_ref_def_func"`
	TimeRefUniqueCheck      pg.NullTime `db:"time_ref_unique_check"`
	TimeDefConst            pg.NullTime `db:"time_def_const"`
	TimeDefConstUniqueCheck pg.NullTime `db:"time_def_const_unique_check"`
	TimeDefFunc             pg.NullTime `db:"time_def_func"`
	TimeDefFuncUniqueCheck  pg.NullTime `db:"time_def_func_unique_check"`
}
