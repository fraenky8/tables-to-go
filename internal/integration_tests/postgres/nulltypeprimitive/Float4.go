package dto

type Float4 struct {
	Float4                    *float64 `db:"float4"`
	Float4Nn                  float64  `db:"float4_nn"`
	Float4NnUnique            float64  `db:"float4_nn_unique"`
	Float4NnCheck             float64  `db:"float4_nn_check"`
	Float4NnRef               float64  `db:"float4_nn_ref"`
	Float4NnDefConst          float64  `db:"float4_nn_def_const"`
	Float4NnDefFunc           float64  `db:"float4_nn_def_func"`
	Float4NnUniqueCheck       float64  `db:"float4_nn_unique_check"`
	Float4Unique              *float64 `db:"float4_unique"`
	Float4UniqueCheck         *float64 `db:"float4_unique_check"`
	Float4UniqueRef           *float64 `db:"float4_unique_ref"`
	Float4UniqueDefConst      *float64 `db:"float4_unique_def_const"`
	Float4UniqueDefFunc       *float64 `db:"float4_unique_def_func"`
	Float4Check               *float64 `db:"float4_check"`
	Float4CheckRef            *float64 `db:"float4_check_ref"`
	Float4CheckDefConst       *float64 `db:"float4_check_def_const"`
	Float4CheckDefFunc        *float64 `db:"float4_check_def_func"`
	Float4Ref                 *float64 `db:"float4_ref"`
	Float4RefDefConst         *float64 `db:"float4_ref_def_const"`
	Float4RefDefFunc          *float64 `db:"float4_ref_def_func"`
	Float4RefUniqueCheck      *float64 `db:"float4_ref_unique_check"`
	Float4DefConst            *float64 `db:"float4_def_const"`
	Float4DefConstUniqueCheck *float64 `db:"float4_def_const_unique_check"`
	Float4DefFunc             *float64 `db:"float4_def_func"`
	Float4DefFuncUniqueCheck  *float64 `db:"float4_def_func_unique_check"`
}
