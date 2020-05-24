package dto

import (
	"database/sql"
)

type TinyintTable struct {
	I                          sql.NullInt64 `db:"i"`
	TinyintNn                  int           `db:"tinyint_nn"`
	TinyintNnUnique            int           `db:"tinyint_nn_unique"`
	TinyintNnCheck             int           `db:"tinyint_nn_check"`
	TinyintUnique              sql.NullInt64 `db:"tinyint_unique"`
	TinyintUniqueCheck         sql.NullInt64 `db:"tinyint_unique_check"`
	TinyintUniqueRef           sql.NullInt64 `db:"tinyint_unique_ref"`
	TinyintUniqueDefConst      sql.NullInt64 `db:"tinyint_unique_def_const"`
	TinyintUniqueDefFunc       sql.NullInt64 `db:"tinyint_unique_def_func"`
	TinyintCheck               sql.NullInt64 `db:"tinyint_check"`
	TinyintCheckRef            sql.NullInt64 `db:"tinyint_check_ref"`
	TinyintCheckDefConst       sql.NullInt64 `db:"tinyint_check_def_const"`
	TinyintCheckDefFunc        sql.NullInt64 `db:"tinyint_check_def_func"`
	TinyintRef                 sql.NullInt64 `db:"tinyint_ref"`
	TinyintRefUniqueCheck      sql.NullInt64 `db:"tinyint_ref_unique_check"`
	TinyintDefConst            sql.NullInt64 `db:"tinyint_def_const"`
	TinyintDefConstUniqueCheck sql.NullInt64 `db:"tinyint_def_const_unique_check"`
	TinyintDefFunc             sql.NullInt64 `db:"tinyint_def_func"`
	TinyintDefFuncUniqueCheck  sql.NullInt64 `db:"tinyint_def_func_unique_check"`
}
