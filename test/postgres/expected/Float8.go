package dto

import (
	"database/sql"
)

type Float8 struct {
	Float8                    sql.NullFloat64 `db:"float8"`
	Float8Nn                  float64         `db:"float8_nn"`
	Float8NnUnique            float64         `db:"float8_nn_unique"`
	Float8NnCheck             float64         `db:"float8_nn_check"`
	Float8NnRef               float64         `db:"float8_nn_ref"`
	Float8NnDefConst          float64         `db:"float8_nn_def_const"`
	Float8NnDefFunc           float64         `db:"float8_nn_def_func"`
	Float8NnUniqueCheck       float64         `db:"float8_nn_unique_check"`
	Float8Unique              sql.NullFloat64 `db:"float8_unique"`
	Float8UniqueCheck         sql.NullFloat64 `db:"float8_unique_check"`
	Float8UniqueRef           sql.NullFloat64 `db:"float8_unique_ref"`
	Float8UniqueDefConst      sql.NullFloat64 `db:"float8_unique_def_const"`
	Float8UniqueDefFunc       sql.NullFloat64 `db:"float8_unique_def_func"`
	Float8Check               sql.NullFloat64 `db:"float8_check"`
	Float8CheckRef            sql.NullFloat64 `db:"float8_check_ref"`
	Float8CheckDefConst       sql.NullFloat64 `db:"float8_check_def_const"`
	Float8CheckDefFunc        sql.NullFloat64 `db:"float8_check_def_func"`
	Float8Ref                 sql.NullFloat64 `db:"float8_ref"`
	Float8RefDefConst         sql.NullFloat64 `db:"float8_ref_def_const"`
	Float8RefDefFunc          sql.NullFloat64 `db:"float8_ref_def_func"`
	Float8RefUniqueCheck      sql.NullFloat64 `db:"float8_ref_unique_check"`
	Float8DefConst            sql.NullFloat64 `db:"float8_def_const"`
	Float8DefConstUniqueCheck sql.NullFloat64 `db:"float8_def_const_unique_check"`
	Float8DefFunc             sql.NullFloat64 `db:"float8_def_func"`
	Float8DefFuncUniqueCheck  sql.NullFloat64 `db:"float8_def_func_unique_check"`
}
