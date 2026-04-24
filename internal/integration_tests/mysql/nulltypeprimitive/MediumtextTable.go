package dto

type MediumtextTable struct {
	Col                     *string `db:"col"`
	MediumtextDefConst      *string `db:"mediumtext_def_const"`
	MediumtextDefFunc       *string `db:"mediumtext_def_func"`
	MediumtextNn            string  `db:"mediumtext_nn"`
	MediumtextNnCheckCmp    string  `db:"mediumtext_nn_check_cmp"`
	MediumtextNnCheckFn     string  `db:"mediumtext_nn_check_fn"`
	MediumtextNnDefConst    string  `db:"mediumtext_nn_def_const"`
	MediumtextNnDefFunc     string  `db:"mediumtext_nn_def_func"`
	MediumtextCheck         *string `db:"mediumtext_check"`
	MediumtextCheckDefConst *string `db:"mediumtext_check_def_const"`
	MediumtextCheckDefFunc  *string `db:"mediumtext_check_def_func"`
}
