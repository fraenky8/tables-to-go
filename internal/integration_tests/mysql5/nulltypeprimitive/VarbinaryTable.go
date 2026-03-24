package dto

type VarbinaryTable struct {
	Col                          *string `db:"col"`
	VarbinaryCap                 *string `db:"varbinary_cap"`
	VarbinaryNn                  string  `db:"varbinary_nn"`
	VarbinaryNnUnique            string  `db:"varbinary_nn_unique"`
	VarbinaryNnCheckCmp          string  `db:"varbinary_nn_check_cmp"`
	VarbinaryNnCheckFn           string  `db:"varbinary_nn_check_fn"`
	VarbinaryNnRef               string  `db:"varbinary_nn_ref"`
	VarbinaryNnDefConst          string  `db:"varbinary_nn_def_const"`
	VarbinaryNnDefFunc           string  `db:"varbinary_nn_def_func"`
	VarbinaryNnUniqueCheck       string  `db:"varbinary_nn_unique_check"`
	VarbinaryUnique              *string `db:"varbinary_unique"`
	VarbinaryUniqueCheck         *string `db:"varbinary_unique_check"`
	VarbinaryUniqueRef           *string `db:"varbinary_unique_ref"`
	VarbinaryUniqueDefConst      *string `db:"varbinary_unique_def_const"`
	VarbinaryUniqueDefFunc       *string `db:"varbinary_unique_def_func"`
	VarbinaryCheck               *string `db:"varbinary_check"`
	VarbinaryCheckRef            *string `db:"varbinary_check_ref"`
	VarbinaryCheckDefConst       *string `db:"varbinary_check_def_const"`
	VarbinaryCheckDefFunc        *string `db:"varbinary_check_def_func"`
	VarbinaryRef                 *string `db:"varbinary_ref"`
	VarbinaryRefUniqueCheck      *string `db:"varbinary_ref_unique_check"`
	VarbinaryDefConst            *string `db:"varbinary_def_const"`
	VarbinaryDefConstUniqueCheck *string `db:"varbinary_def_const_unique_check"`
	VarbinaryDefFunc             *string `db:"varbinary_def_func"`
	VarbinaryDefFuncUniqueCheck  *string `db:"varbinary_def_func_unique_check"`
}
