package dto

type BigintTable struct {
	I                         *int `db:"i"`
	BigintNn                  int  `db:"bigint_nn"`
	BigintNnUnique            int  `db:"bigint_nn_unique"`
	BigintNnCheck             int  `db:"bigint_nn_check"`
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
	BigintRefUniqueCheck      *int `db:"bigint_ref_unique_check"`
	BigintDefConst            *int `db:"bigint_def_const"`
	BigintDefConstUniqueCheck *int `db:"bigint_def_const_unique_check"`
	BigintDefFunc             *int `db:"bigint_def_func"`
	BigintDefFuncUniqueCheck  *int `db:"bigint_def_func_unique_check"`
}
