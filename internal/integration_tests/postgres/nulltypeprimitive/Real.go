package dto

type Real struct {
	Real                    *float64 `db:"real"`
	RealNn                  float64  `db:"real_nn"`
	RealNnUnique            float64  `db:"real_nn_unique"`
	RealNnCheck             float64  `db:"real_nn_check"`
	RealNnRef               float64  `db:"real_nn_ref"`
	RealNnDefConst          float64  `db:"real_nn_def_const"`
	RealNnDefFunc           float64  `db:"real_nn_def_func"`
	RealNnUniqueCheck       float64  `db:"real_nn_unique_check"`
	RealUnique              *float64 `db:"real_unique"`
	RealUniqueCheck         *float64 `db:"real_unique_check"`
	RealUniqueRef           *float64 `db:"real_unique_ref"`
	RealUniqueDefConst      *float64 `db:"real_unique_def_const"`
	RealUniqueDefFunc       *float64 `db:"real_unique_def_func"`
	RealCheck               *float64 `db:"real_check"`
	RealCheckRef            *float64 `db:"real_check_ref"`
	RealCheckDefConst       *float64 `db:"real_check_def_const"`
	RealCheckDefFunc        *float64 `db:"real_check_def_func"`
	RealRef                 *float64 `db:"real_ref"`
	RealRefDefConst         *float64 `db:"real_ref_def_const"`
	RealRefDefFunc          *float64 `db:"real_ref_def_func"`
	RealRefUniqueCheck      *float64 `db:"real_ref_unique_check"`
	RealDefConst            *float64 `db:"real_def_const"`
	RealDefConstUniqueCheck *float64 `db:"real_def_const_unique_check"`
	RealDefFunc             *float64 `db:"real_def_func"`
	RealDefFuncUniqueCheck  *float64 `db:"real_def_func_unique_check"`
}
