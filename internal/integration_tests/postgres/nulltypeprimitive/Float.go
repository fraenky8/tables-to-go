package dto

type Float struct {
	Float                    *float64 `db:"float"`
	FloatNn                  float64  `db:"float_nn"`
	FloatNnUnique            float64  `db:"float_nn_unique"`
	FloatNnCheck             float64  `db:"float_nn_check"`
	FloatNnRef               float64  `db:"float_nn_ref"`
	FloatNnDefConst          float64  `db:"float_nn_def_const"`
	FloatNnDefFunc           float64  `db:"float_nn_def_func"`
	FloatNnUniqueCheck       float64  `db:"float_nn_unique_check"`
	FloatUnique              *float64 `db:"float_unique"`
	FloatUniqueCheck         *float64 `db:"float_unique_check"`
	FloatUniqueRef           *float64 `db:"float_unique_ref"`
	FloatUniqueDefConst      *float64 `db:"float_unique_def_const"`
	FloatUniqueDefFunc       *float64 `db:"float_unique_def_func"`
	FloatCheck               *float64 `db:"float_check"`
	FloatCheckRef            *float64 `db:"float_check_ref"`
	FloatCheckDefConst       *float64 `db:"float_check_def_const"`
	FloatCheckDefFunc        *float64 `db:"float_check_def_func"`
	FloatRef                 *float64 `db:"float_ref"`
	FloatRefDefConst         *float64 `db:"float_ref_def_const"`
	FloatRefDefFunc          *float64 `db:"float_ref_def_func"`
	FloatRefUniqueCheck      *float64 `db:"float_ref_unique_check"`
	FloatDefConst            *float64 `db:"float_def_const"`
	FloatDefConstUniqueCheck *float64 `db:"float_def_const_unique_check"`
	FloatDefFunc             *float64 `db:"float_def_func"`
	FloatDefFuncUniqueCheck  *float64 `db:"float_def_func_unique_check"`
}
