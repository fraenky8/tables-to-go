package dto

type TinyblobTable struct {
	Col                   *string `db:"col"`
	TinyblobDefConst      *string `db:"tinyblob_def_const"`
	TinyblobDefFunc       *string `db:"tinyblob_def_func"`
	TinyblobNn            string  `db:"tinyblob_nn"`
	TinyblobNnCheckCmp    string  `db:"tinyblob_nn_check_cmp"`
	TinyblobNnCheckFn     string  `db:"tinyblob_nn_check_fn"`
	TinyblobNnDefConst    string  `db:"tinyblob_nn_def_const"`
	TinyblobNnDefFunc     string  `db:"tinyblob_nn_def_func"`
	TinyblobCheck         *string `db:"tinyblob_check"`
	TinyblobCheckDefConst *string `db:"tinyblob_check_def_const"`
	TinyblobCheckDefFunc  *string `db:"tinyblob_check_def_func"`
}
