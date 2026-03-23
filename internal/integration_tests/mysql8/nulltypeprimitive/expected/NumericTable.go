package dto

type NumericTable struct {
	Col                        *float64 `db:"col"`
	NumericNn                  float64  `db:"numeric_nn"`
	NumericNnUnique            float64  `db:"numeric_nn_unique"`
	NumericNnCheck             float64  `db:"numeric_nn_check"`
	NumericNnRef               float64  `db:"numeric_nn_ref"`
	NumericNnDefConst          float64  `db:"numeric_nn_def_const"`
	NumericNnDefFunc           float64  `db:"numeric_nn_def_func"`
	NumericNnUniqueCheck       float64  `db:"numeric_nn_unique_check"`
	NumericUnique              *float64 `db:"numeric_unique"`
	NumericUniqueCheck         *float64 `db:"numeric_unique_check"`
	NumericUniqueRef           *float64 `db:"numeric_unique_ref"`
	NumericUniqueDefConst      *float64 `db:"numeric_unique_def_const"`
	NumericUniqueDefFunc       *float64 `db:"numeric_unique_def_func"`
	NumericCheck               *float64 `db:"numeric_check"`
	NumericCheckRef            *float64 `db:"numeric_check_ref"`
	NumericCheckDefConst       *float64 `db:"numeric_check_def_const"`
	NumericCheckDefFunc        *float64 `db:"numeric_check_def_func"`
	NumericRef                 *float64 `db:"numeric_ref"`
	NumericRefUniqueCheck      *float64 `db:"numeric_ref_unique_check"`
	NumericDefConst            *float64 `db:"numeric_def_const"`
	NumericDefConstUniqueCheck *float64 `db:"numeric_def_const_unique_check"`
	NumericDefFunc             *float64 `db:"numeric_def_func"`
	NumericDefFuncUniqueCheck  *float64 `db:"numeric_def_func_unique_check"`
}
