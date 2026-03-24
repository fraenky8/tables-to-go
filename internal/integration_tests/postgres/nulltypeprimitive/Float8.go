package dto

type Float8 struct {
	Float8                    *float64 `db:"float8"`
	Float8Nn                  float64  `db:"float8_nn"`
	Float8NnUnique            float64  `db:"float8_nn_unique"`
	Float8NnCheck             float64  `db:"float8_nn_check"`
	Float8NnRef               float64  `db:"float8_nn_ref"`
	Float8NnDefConst          float64  `db:"float8_nn_def_const"`
	Float8NnDefFunc           float64  `db:"float8_nn_def_func"`
	Float8NnUniqueCheck       float64  `db:"float8_nn_unique_check"`
	Float8Unique              *float64 `db:"float8_unique"`
	Float8UniqueCheck         *float64 `db:"float8_unique_check"`
	Float8UniqueRef           *float64 `db:"float8_unique_ref"`
	Float8UniqueDefConst      *float64 `db:"float8_unique_def_const"`
	Float8UniqueDefFunc       *float64 `db:"float8_unique_def_func"`
	Float8Check               *float64 `db:"float8_check"`
	Float8CheckRef            *float64 `db:"float8_check_ref"`
	Float8CheckDefConst       *float64 `db:"float8_check_def_const"`
	Float8CheckDefFunc        *float64 `db:"float8_check_def_func"`
	Float8Ref                 *float64 `db:"float8_ref"`
	Float8RefDefConst         *float64 `db:"float8_ref_def_const"`
	Float8RefDefFunc          *float64 `db:"float8_ref_def_func"`
	Float8RefUniqueCheck      *float64 `db:"float8_ref_unique_check"`
	Float8DefConst            *float64 `db:"float8_def_const"`
	Float8DefConstUniqueCheck *float64 `db:"float8_def_const_unique_check"`
	Float8DefFunc             *float64 `db:"float8_def_func"`
	Float8DefFuncUniqueCheck  *float64 `db:"float8_def_func_unique_check"`
}
