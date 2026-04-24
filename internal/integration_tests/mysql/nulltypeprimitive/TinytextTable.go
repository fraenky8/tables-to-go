package dto

type TinytextTable struct {
	Col                   *string `db:"col"`
	TinytextDefConst      *string `db:"tinytext_def_const"`
	TinytextDefFunc       *string `db:"tinytext_def_func"`
	TinytextNn            string  `db:"tinytext_nn"`
	TinytextNnCheckCmp    string  `db:"tinytext_nn_check_cmp"`
	TinytextNnCheckFn     string  `db:"tinytext_nn_check_fn"`
	TinytextNnDefConst    string  `db:"tinytext_nn_def_const"`
	TinytextNnDefFunc     string  `db:"tinytext_nn_def_func"`
	TinytextCheck         *string `db:"tinytext_check"`
	TinytextCheckDefConst *string `db:"tinytext_check_def_const"`
	TinytextCheckDefFunc  *string `db:"tinytext_check_def_func"`
}
