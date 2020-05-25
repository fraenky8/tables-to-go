package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type Date struct {
	Date                    pg.NullTime `db:"date"`
	DateNn                  time.Time   `db:"date_nn"`
	DateNnUnique            time.Time   `db:"date_nn_unique"`
	DateNnCheck             time.Time   `db:"date_nn_check"`
	DateNnRef               time.Time   `db:"date_nn_ref"`
	DateNnDefConst          time.Time   `db:"date_nn_def_const"`
	DateNnDefFunc           time.Time   `db:"date_nn_def_func"`
	DateNnUniqueCheck       time.Time   `db:"date_nn_unique_check"`
	DateUnique              pg.NullTime `db:"date_unique"`
	DateUniqueCheck         pg.NullTime `db:"date_unique_check"`
	DateUniqueRef           pg.NullTime `db:"date_unique_ref"`
	DateUniqueDefConst      pg.NullTime `db:"date_unique_def_const"`
	DateUniqueDefFunc       pg.NullTime `db:"date_unique_def_func"`
	DateCheck               pg.NullTime `db:"date_check"`
	DateCheckRef            pg.NullTime `db:"date_check_ref"`
	DateCheckDefConst       pg.NullTime `db:"date_check_def_const"`
	DateCheckDefFunc        pg.NullTime `db:"date_check_def_func"`
	DateRef                 pg.NullTime `db:"date_ref"`
	DateRefDefConst         pg.NullTime `db:"date_ref_def_const"`
	DateRefDefFunc          pg.NullTime `db:"date_ref_def_func"`
	DateRefUniqueCheck      pg.NullTime `db:"date_ref_unique_check"`
	DateDefConst            pg.NullTime `db:"date_def_const"`
	DateDefConstUniqueCheck pg.NullTime `db:"date_def_const_unique_check"`
	DateDefFunc             pg.NullTime `db:"date_def_func"`
	DateDefFuncUniqueCheck  pg.NullTime `db:"date_def_func_unique_check"`
}
