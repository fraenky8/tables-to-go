package dto

type VarcharTable struct {
	Col                        *string `db:"col"`
	VarcharCap                 *string `db:"varchar_cap"`
	VarcharNn                  string  `db:"varchar_nn"`
	VarcharNnUnique            string  `db:"varchar_nn_unique"`
	VarcharNnCheckCmp          string  `db:"varchar_nn_check_cmp"`
	VarcharNnCheckFn           string  `db:"varchar_nn_check_fn"`
	VarcharNnRef               string  `db:"varchar_nn_ref"`
	VarcharNnDefConst          string  `db:"varchar_nn_def_const"`
	VarcharNnDefFunc           string  `db:"varchar_nn_def_func"`
	VarcharNnUniqueCheck       string  `db:"varchar_nn_unique_check"`
	VarcharUnique              *string `db:"varchar_unique"`
	VarcharUniqueCheck         *string `db:"varchar_unique_check"`
	VarcharUniqueRef           *string `db:"varchar_unique_ref"`
	VarcharUniqueDefConst      *string `db:"varchar_unique_def_const"`
	VarcharUniqueDefFunc       *string `db:"varchar_unique_def_func"`
	VarcharCheck               *string `db:"varchar_check"`
	VarcharCheckRef            *string `db:"varchar_check_ref"`
	VarcharCheckDefConst       *string `db:"varchar_check_def_const"`
	VarcharCheckDefFunc        *string `db:"varchar_check_def_func"`
	VarcharRef                 *string `db:"varchar_ref"`
	VarcharRefUniqueCheck      *string `db:"varchar_ref_unique_check"`
	VarcharDefConst            *string `db:"varchar_def_const"`
	VarcharDefConstUniqueCheck *string `db:"varchar_def_const_unique_check"`
	VarcharDefFunc             *string `db:"varchar_def_func"`
	VarcharDefFuncUniqueCheck  *string `db:"varchar_def_func_unique_check"`
}
