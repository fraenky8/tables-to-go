package dto

type DecimalTable struct {
	Col                        *float64 `db:"col"`
	DecimalNn                  float64  `db:"decimal_nn"`
	DecimalNnUnique            float64  `db:"decimal_nn_unique"`
	DecimalNnCheck             float64  `db:"decimal_nn_check"`
	DecimalNnRef               float64  `db:"decimal_nn_ref"`
	DecimalNnDefConst          float64  `db:"decimal_nn_def_const"`
	DecimalNnDefFunc           float64  `db:"decimal_nn_def_func"`
	DecimalNnUniqueCheck       float64  `db:"decimal_nn_unique_check"`
	DecimalUnique              *float64 `db:"decimal_unique"`
	DecimalUniqueCheck         *float64 `db:"decimal_unique_check"`
	DecimalUniqueRef           *float64 `db:"decimal_unique_ref"`
	DecimalUniqueDefConst      *float64 `db:"decimal_unique_def_const"`
	DecimalUniqueDefFunc       *float64 `db:"decimal_unique_def_func"`
	DecimalCheck               *float64 `db:"decimal_check"`
	DecimalCheckRef            *float64 `db:"decimal_check_ref"`
	DecimalCheckDefConst       *float64 `db:"decimal_check_def_const"`
	DecimalCheckDefFunc        *float64 `db:"decimal_check_def_func"`
	DecimalRef                 *float64 `db:"decimal_ref"`
	DecimalRefUniqueCheck      *float64 `db:"decimal_ref_unique_check"`
	DecimalDefConst            *float64 `db:"decimal_def_const"`
	DecimalDefConstUniqueCheck *float64 `db:"decimal_def_const_unique_check"`
	DecimalDefFunc             *float64 `db:"decimal_def_func"`
	DecimalDefFuncUniqueCheck  *float64 `db:"decimal_def_func_unique_check"`
}
