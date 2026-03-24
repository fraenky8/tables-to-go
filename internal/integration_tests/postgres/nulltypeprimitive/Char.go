package dto

type Char struct {
	Char                    *string `db:"char"`
	CharCap                 *string `db:"char_cap"`
	CharNn                  string  `db:"char_nn"`
	CharNnUnique            string  `db:"char_nn_unique"`
	CharNnCheckCmp          string  `db:"char_nn_check_cmp"`
	CharNnCheckFn           string  `db:"char_nn_check_fn"`
	CharNnRef               string  `db:"char_nn_ref"`
	CharNnDefConst          string  `db:"char_nn_def_const"`
	CharNnDefFunc           string  `db:"char_nn_def_func"`
	CharNnUniqueCheck       string  `db:"char_nn_unique_check"`
	CharUnique              *string `db:"char_unique"`
	CharUniqueCheck         *string `db:"char_unique_check"`
	CharUniqueRef           *string `db:"char_unique_ref"`
	CharUniqueDefConst      *string `db:"char_unique_def_const"`
	CharUniqueDefFunc       *string `db:"char_unique_def_func"`
	CharCheck               *string `db:"char_check"`
	CharCheckRef            *string `db:"char_check_ref"`
	CharCheckDefConst       *string `db:"char_check_def_const"`
	CharCheckDefFunc        *string `db:"char_check_def_func"`
	CharRef                 *string `db:"char_ref"`
	CharRefDefConst         *string `db:"char_ref_def_const"`
	CharRefDefFunc          *string `db:"char_ref_def_func"`
	CharRefUniqueCheck      *string `db:"char_ref_unique_check"`
	CharDefConst            *string `db:"char_def_const"`
	CharDefConstUniqueCheck *string `db:"char_def_const_unique_check"`
	CharDefFunc             *string `db:"char_def_func"`
	CharDefFuncUniqueCheck  *string `db:"char_def_func_unique_check"`
}
