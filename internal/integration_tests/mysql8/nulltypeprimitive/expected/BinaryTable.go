package dto

type BinaryTable struct {
	Col                       *string `db:"col"`
	BinaryCap                 *string `db:"binary_cap"`
	BinaryNn                  string  `db:"binary_nn"`
	BinaryNnUnique            string  `db:"binary_nn_unique"`
	BinaryNnCheckCmp          string  `db:"binary_nn_check_cmp"`
	BinaryNnCheckFn           string  `db:"binary_nn_check_fn"`
	BinaryNnRef               string  `db:"binary_nn_ref"`
	BinaryNnDefConst          string  `db:"binary_nn_def_const"`
	BinaryNnDefFunc           string  `db:"binary_nn_def_func"`
	BinaryNnUniqueCheck       string  `db:"binary_nn_unique_check"`
	BinaryUnique              *string `db:"binary_unique"`
	BinaryUniqueCheck         *string `db:"binary_unique_check"`
	BinaryUniqueRef           *string `db:"binary_unique_ref"`
	BinaryUniqueDefConst      *string `db:"binary_unique_def_const"`
	BinaryUniqueDefFunc       *string `db:"binary_unique_def_func"`
	BinaryCheck               *string `db:"binary_check"`
	BinaryCheckRef            *string `db:"binary_check_ref"`
	BinaryCheckDefConst       *string `db:"binary_check_def_const"`
	BinaryCheckDefFunc        *string `db:"binary_check_def_func"`
	BinaryRef                 *string `db:"binary_ref"`
	BinaryRefUniqueCheck      *string `db:"binary_ref_unique_check"`
	BinaryDefConst            *string `db:"binary_def_const"`
	BinaryDefConstUniqueCheck *string `db:"binary_def_const_unique_check"`
	BinaryDefFunc             *string `db:"binary_def_func"`
	BinaryDefFuncUniqueCheck  *string `db:"binary_def_func_unique_check"`
}
