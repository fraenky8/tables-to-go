package dto

type LongblobTable struct {
	Col                   *string `db:"col"`
	LongblobDefConst      *string `db:"longblob_def_const"`
	LongblobDefFunc       *string `db:"longblob_def_func"`
	LongblobNn            string  `db:"longblob_nn"`
	LongblobNnCheckCmp    string  `db:"longblob_nn_check_cmp"`
	LongblobNnCheckFn     string  `db:"longblob_nn_check_fn"`
	LongblobNnDefConst    string  `db:"longblob_nn_def_const"`
	LongblobNnDefFunc     string  `db:"longblob_nn_def_func"`
	LongblobCheck         *string `db:"longblob_check"`
	LongblobCheckDefConst *string `db:"longblob_check_def_const"`
	LongblobCheckDefFunc  *string `db:"longblob_check_def_func"`
}
