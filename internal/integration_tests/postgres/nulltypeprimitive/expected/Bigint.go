package dto

type Bigint struct {
	Bigint                    *int `db:"bigint"`
	BigintNn                  int  `db:"bigint_nn"`
	BigintNnUnique            int  `db:"bigint_nn_unique"`
	BigintNnCheck             int  `db:"bigint_nn_check"`
	BigintNnRef               int  `db:"bigint_nn_ref"`
	BigintNnDefConst          int  `db:"bigint_nn_def_const"`
	BigintNnDefFunc           int  `db:"bigint_nn_def_func"`
	BigintNnUniqueCheck       int  `db:"bigint_nn_unique_check"`
	BigintUnique              *int `db:"bigint_unique"`
	BigintUniqueCheck         *int `db:"bigint_unique_check"`
	BigintUniqueRef           *int `db:"bigint_unique_ref"`
	BigintUniqueDefConst      *int `db:"bigint_unique_def_const"`
	BigintUniqueDefFunc       *int `db:"bigint_unique_def_func"`
	BigintCheck               *int `db:"bigint_check"`
	BigintCheckRef            *int `db:"bigint_check_ref"`
	BigintCheckDefConst       *int `db:"bigint_check_def_const"`
	BigintCheckDefFunc        *int `db:"bigint_check_def_func"`
	BigintRef                 *int `db:"bigint_ref"`
	BigintRefDefConst         *int `db:"bigint_ref_def_const"`
	BigintRefDefFunc          *int `db:"bigint_ref_def_func"`
	BigintRefUniqueCheck      *int `db:"bigint_ref_unique_check"`
	BigintDefConst            *int `db:"bigint_def_const"`
	BigintDefConstUniqueCheck *int `db:"bigint_def_const_unique_check"`
	BigintDefFunc             *int `db:"bigint_def_func"`
	BigintDefFuncUniqueCheck  *int `db:"bigint_def_func_unique_check"`
}
