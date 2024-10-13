package dto

import (
	"database/sql"
	"time"
)

type YearTable struct {
	Year                    sql.NullTime `db:"year"`
	YearNn                  time.Time    `db:"year_nn"`
	YearNnUnique            time.Time    `db:"year_nn_unique"`
	YearNnCheck             time.Time    `db:"year_nn_check"`
	YearNnRef               time.Time    `db:"year_nn_ref"`
	YearNnDefConst          time.Time    `db:"year_nn_def_const"`
	YearNnDefFunc           time.Time    `db:"year_nn_def_func"`
	YearNnUniqueCheck       time.Time    `db:"year_nn_unique_check"`
	YearUnique              sql.NullTime `db:"year_unique"`
	YearUniqueCheck         sql.NullTime `db:"year_unique_check"`
	YearUniqueRef           sql.NullTime `db:"year_unique_ref"`
	YearUniqueDefConst      sql.NullTime `db:"year_unique_def_const"`
	YearUniqueDefFunc       sql.NullTime `db:"year_unique_def_func"`
	YearCheck               sql.NullTime `db:"year_check"`
	YearCheckRef            sql.NullTime `db:"year_check_ref"`
	YearCheckDefConst       sql.NullTime `db:"year_check_def_const"`
	YearCheckDefFunc        sql.NullTime `db:"year_check_def_func"`
	YearRef                 sql.NullTime `db:"year_ref"`
	YearRefUniqueCheck      sql.NullTime `db:"year_ref_unique_check"`
	YearDefConst            sql.NullTime `db:"year_def_const"`
	YearDefConstUniqueCheck sql.NullTime `db:"year_def_const_unique_check"`
	YearDefFunc             sql.NullTime `db:"year_def_func"`
	YearDefFuncUniqueCheck  sql.NullTime `db:"year_def_func_unique_check"`
}
