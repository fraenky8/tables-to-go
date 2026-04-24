package dto

import (
	"database/sql"
)

type FloatTable struct {
	Col                      sql.NullFloat64 `db:"col"`
	FloatNn                  float64         `db:"float_nn"`
	FloatNnUnique            float64         `db:"float_nn_unique"`
	FloatNnCheck             float64         `db:"float_nn_check"`
	FloatNnRef               float64         `db:"float_nn_ref"`
	FloatNnDefConst          float64         `db:"float_nn_def_const"`
	FloatNnDefFunc           float64         `db:"float_nn_def_func"`
	FloatNnUniqueCheck       float64         `db:"float_nn_unique_check"`
	FloatUnique              sql.NullFloat64 `db:"float_unique"`
	FloatUniqueCheck         sql.NullFloat64 `db:"float_unique_check"`
	FloatUniqueRef           sql.NullFloat64 `db:"float_unique_ref"`
	FloatUniqueDefConst      sql.NullFloat64 `db:"float_unique_def_const"`
	FloatUniqueDefFunc       sql.NullFloat64 `db:"float_unique_def_func"`
	FloatCheck               sql.NullFloat64 `db:"float_check"`
	FloatCheckRef            sql.NullFloat64 `db:"float_check_ref"`
	FloatCheckDefConst       sql.NullFloat64 `db:"float_check_def_const"`
	FloatCheckDefFunc        sql.NullFloat64 `db:"float_check_def_func"`
	FloatRef                 sql.NullFloat64 `db:"float_ref"`
	FloatRefUniqueCheck      sql.NullFloat64 `db:"float_ref_unique_check"`
	FloatDefConst            sql.NullFloat64 `db:"float_def_const"`
	FloatDefConstUniqueCheck sql.NullFloat64 `db:"float_def_const_unique_check"`
	FloatDefFunc             sql.NullFloat64 `db:"float_def_func"`
	FloatDefFuncUniqueCheck  sql.NullFloat64 `db:"float_def_func_unique_check"`
}
