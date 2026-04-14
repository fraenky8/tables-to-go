package dto

type TinyblobTable struct {
	Col                   *string `db:"col"`
	TinyblobDefConst      *string `db:"tinyblob_def_const"`
	TinyblobDefFunc       *string `db:"tinyblob_def_func"`
	TinyblobRef           *string `db:"tinyblob_ref"`
	TinyblobNn            string  `db:"tinyblob_nn"`
	TinyblobNnCheckCmp    string  `db:"tinyblob_nn_check_cmp"`
	TinyblobNnCheckFn     string  `db:"tinyblob_nn_check_fn"`
	TinyblobNnRef         string  `db:"tinyblob_nn_ref"`
	TinyblobNnDefConst    string  `db:"tinyblob_nn_def_const"`
	TinyblobNnDefFunc     string  `db:"tinyblob_nn_def_func"`
	TinyblobCheck         *string `db:"tinyblob_check"`
	TinyblobCheckRef      *string `db:"tinyblob_check_ref"`
	TinyblobCheckDefConst *string `db:"tinyblob_check_def_const"`
	TinyblobCheckDefFunc  *string `db:"tinyblob_check_def_func"`
}
