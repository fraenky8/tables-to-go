package dto

import (
	"database/sql"
)

type Smallint struct {
	Smallint                    sql.NullInt64 `db:"smallint"`
	SmallintNn                  int           `db:"smallint_nn"`
	SmallintNnUnique            int           `db:"smallint_nn_unique"`
	SmallintNnCheck             int           `db:"smallint_nn_check"`
	SmallintNnRef               int           `db:"smallint_nn_ref"`
	SmallintNnDefConst          int           `db:"smallint_nn_def_const"`
	SmallintNnDefFunc           int           `db:"smallint_nn_def_func"`
	SmallintNnUniqueCheck       int           `db:"smallint_nn_unique_check"`
	SmallintUnique              sql.NullInt64 `db:"smallint_unique"`
	SmallintUniqueCheck         sql.NullInt64 `db:"smallint_unique_check"`
	SmallintUniqueRef           sql.NullInt64 `db:"smallint_unique_ref"`
	SmallintUniqueDefConst      sql.NullInt64 `db:"smallint_unique_def_const"`
	SmallintUniqueDefFunc       sql.NullInt64 `db:"smallint_unique_def_func"`
	SmallintCheck               sql.NullInt64 `db:"smallint_check"`
	SmallintCheckRef            sql.NullInt64 `db:"smallint_check_ref"`
	SmallintCheckDefConst       sql.NullInt64 `db:"smallint_check_def_const"`
	SmallintCheckDefFunc        sql.NullInt64 `db:"smallint_check_def_func"`
	SmallintRef                 sql.NullInt64 `db:"smallint_ref"`
	SmallintRefDefConst         sql.NullInt64 `db:"smallint_ref_def_const"`
	SmallintRefDefFunc          sql.NullInt64 `db:"smallint_ref_def_func"`
	SmallintRefUniqueCheck      sql.NullInt64 `db:"smallint_ref_unique_check"`
	SmallintDefConst            sql.NullInt64 `db:"smallint_def_const"`
	SmallintDefConstUniqueCheck sql.NullInt64 `db:"smallint_def_const_unique_check"`
	SmallintDefFunc             sql.NullInt64 `db:"smallint_def_func"`
	SmallintDefFuncUniqueCheck  sql.NullInt64 `db:"smallint_def_func_unique_check"`
}
