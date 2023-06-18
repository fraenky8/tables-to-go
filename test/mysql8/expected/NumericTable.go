package dto

import (
	"database/sql"
)

type NumericTable struct {
	Col                        sql.NullFloat64 `db:"col"`
	NumericNn                  float64         `db:"numeric_nn"`
	NumericNnUnique            float64         `db:"numeric_nn_unique"`
	NumericNnCheck             float64         `db:"numeric_nn_check"`
	NumericNnRef               float64         `db:"numeric_nn_ref"`
	NumericNnDefConst          float64         `db:"numeric_nn_def_const"`
	NumericNnDefFunc           float64         `db:"numeric_nn_def_func"`
	NumericNnUniqueCheck       float64         `db:"numeric_nn_unique_check"`
	NumericUnique              sql.NullFloat64 `db:"numeric_unique"`
	NumericUniqueCheck         sql.NullFloat64 `db:"numeric_unique_check"`
	NumericUniqueRef           sql.NullFloat64 `db:"numeric_unique_ref"`
	NumericUniqueDefConst      sql.NullFloat64 `db:"numeric_unique_def_const"`
	NumericUniqueDefFunc       sql.NullFloat64 `db:"numeric_unique_def_func"`
	NumericCheck               sql.NullFloat64 `db:"numeric_check"`
	NumericCheckRef            sql.NullFloat64 `db:"numeric_check_ref"`
	NumericCheckDefConst       sql.NullFloat64 `db:"numeric_check_def_const"`
	NumericCheckDefFunc        sql.NullFloat64 `db:"numeric_check_def_func"`
	NumericRef                 sql.NullFloat64 `db:"numeric_ref"`
	NumericRefUniqueCheck      sql.NullFloat64 `db:"numeric_ref_unique_check"`
	NumericDefConst            sql.NullFloat64 `db:"numeric_def_const"`
	NumericDefConstUniqueCheck sql.NullFloat64 `db:"numeric_def_const_unique_check"`
	NumericDefFunc             sql.NullFloat64 `db:"numeric_def_func"`
	NumericDefFuncUniqueCheck  sql.NullFloat64 `db:"numeric_def_func_unique_check"`
}
