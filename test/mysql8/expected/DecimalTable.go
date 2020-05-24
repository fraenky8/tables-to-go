package dto

import (
	"database/sql"
)

type DecimalTable struct {
	Col                        sql.NullFloat64 `db:"col"`
	DecimalNn                  float64         `db:"decimal_nn"`
	DecimalNnUnique            float64         `db:"decimal_nn_unique"`
	DecimalNnCheck             float64         `db:"decimal_nn_check"`
	DecimalNnRef               float64         `db:"decimal_nn_ref"`
	DecimalNnDefConst          float64         `db:"decimal_nn_def_const"`
	DecimalNnDefFunc           float64         `db:"decimal_nn_def_func"`
	DecimalNnUniqueCheck       float64         `db:"decimal_nn_unique_check"`
	DecimalUnique              sql.NullFloat64 `db:"decimal_unique"`
	DecimalUniqueCheck         sql.NullFloat64 `db:"decimal_unique_check"`
	DecimalUniqueRef           sql.NullFloat64 `db:"decimal_unique_ref"`
	DecimalUniqueDefConst      sql.NullFloat64 `db:"decimal_unique_def_const"`
	DecimalUniqueDefFunc       sql.NullFloat64 `db:"decimal_unique_def_func"`
	DecimalCheck               sql.NullFloat64 `db:"decimal_check"`
	DecimalCheckRef            sql.NullFloat64 `db:"decimal_check_ref"`
	DecimalCheckDefConst       sql.NullFloat64 `db:"decimal_check_def_const"`
	DecimalCheckDefFunc        sql.NullFloat64 `db:"decimal_check_def_func"`
	DecimalRef                 sql.NullFloat64 `db:"decimal_ref"`
	DecimalRefUniqueCheck      sql.NullFloat64 `db:"decimal_ref_unique_check"`
	DecimalDefConst            sql.NullFloat64 `db:"decimal_def_const"`
	DecimalDefConstUniqueCheck sql.NullFloat64 `db:"decimal_def_const_unique_check"`
	DecimalDefFunc             sql.NullFloat64 `db:"decimal_def_func"`
	DecimalDefFuncUniqueCheck  sql.NullFloat64 `db:"decimal_def_func_unique_check"`
}
