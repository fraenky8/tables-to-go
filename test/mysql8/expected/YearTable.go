package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type YearTable struct {
	Year                    mysql.NullTime `db:"year"`
	YearNn                  time.Time      `db:"year_nn"`
	YearNnUnique            time.Time      `db:"year_nn_unique"`
	YearNnCheck             time.Time      `db:"year_nn_check"`
	YearNnRef               time.Time      `db:"year_nn_ref"`
	YearNnDefConst          time.Time      `db:"year_nn_def_const"`
	YearNnDefFunc           time.Time      `db:"year_nn_def_func"`
	YearNnUniqueCheck       time.Time      `db:"year_nn_unique_check"`
	YearUnique              mysql.NullTime `db:"year_unique"`
	YearUniqueCheck         mysql.NullTime `db:"year_unique_check"`
	YearUniqueRef           mysql.NullTime `db:"year_unique_ref"`
	YearUniqueDefConst      mysql.NullTime `db:"year_unique_def_const"`
	YearUniqueDefFunc       mysql.NullTime `db:"year_unique_def_func"`
	YearCheck               mysql.NullTime `db:"year_check"`
	YearCheckRef            mysql.NullTime `db:"year_check_ref"`
	YearCheckDefConst       mysql.NullTime `db:"year_check_def_const"`
	YearCheckDefFunc        mysql.NullTime `db:"year_check_def_func"`
	YearRef                 mysql.NullTime `db:"year_ref"`
	YearRefUniqueCheck      mysql.NullTime `db:"year_ref_unique_check"`
	YearDefConst            mysql.NullTime `db:"year_def_const"`
	YearDefConstUniqueCheck mysql.NullTime `db:"year_def_const_unique_check"`
	YearDefFunc             mysql.NullTime `db:"year_def_func"`
	YearDefFuncUniqueCheck  mysql.NullTime `db:"year_def_func_unique_check"`
}
