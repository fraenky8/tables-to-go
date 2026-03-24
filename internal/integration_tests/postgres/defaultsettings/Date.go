package dto

import (
	"database/sql"
	"time"
)

type Date struct {
	Date                    sql.NullTime `db:"date"`
	DateNn                  time.Time    `db:"date_nn"`
	DateNnUnique            time.Time    `db:"date_nn_unique"`
	DateNnCheck             time.Time    `db:"date_nn_check"`
	DateNnRef               time.Time    `db:"date_nn_ref"`
	DateNnDefConst          time.Time    `db:"date_nn_def_const"`
	DateNnDefFunc           time.Time    `db:"date_nn_def_func"`
	DateNnUniqueCheck       time.Time    `db:"date_nn_unique_check"`
	DateUnique              sql.NullTime `db:"date_unique"`
	DateUniqueCheck         sql.NullTime `db:"date_unique_check"`
	DateUniqueRef           sql.NullTime `db:"date_unique_ref"`
	DateUniqueDefConst      sql.NullTime `db:"date_unique_def_const"`
	DateUniqueDefFunc       sql.NullTime `db:"date_unique_def_func"`
	DateCheck               sql.NullTime `db:"date_check"`
	DateCheckRef            sql.NullTime `db:"date_check_ref"`
	DateCheckDefConst       sql.NullTime `db:"date_check_def_const"`
	DateCheckDefFunc        sql.NullTime `db:"date_check_def_func"`
	DateRef                 sql.NullTime `db:"date_ref"`
	DateRefDefConst         sql.NullTime `db:"date_ref_def_const"`
	DateRefDefFunc          sql.NullTime `db:"date_ref_def_func"`
	DateRefUniqueCheck      sql.NullTime `db:"date_ref_unique_check"`
	DateDefConst            sql.NullTime `db:"date_def_const"`
	DateDefConstUniqueCheck sql.NullTime `db:"date_def_const_unique_check"`
	DateDefFunc             sql.NullTime `db:"date_def_func"`
	DateDefFuncUniqueCheck  sql.NullTime `db:"date_def_func_unique_check"`
}
