package dto

type MediumblobTable struct {
	Col                     *string `db:"col"`
	MediumblobDefConst      *string `db:"mediumblob_def_const"`
	MediumblobDefFunc       *string `db:"mediumblob_def_func"`
	MediumblobNn            string  `db:"mediumblob_nn"`
	MediumblobNnCheckCmp    string  `db:"mediumblob_nn_check_cmp"`
	MediumblobNnCheckFn     string  `db:"mediumblob_nn_check_fn"`
	MediumblobNnDefConst    string  `db:"mediumblob_nn_def_const"`
	MediumblobNnDefFunc     string  `db:"mediumblob_nn_def_func"`
	MediumblobCheck         *string `db:"mediumblob_check"`
	MediumblobCheckDefConst *string `db:"mediumblob_check_def_const"`
	MediumblobCheckDefFunc  *string `db:"mediumblob_check_def_func"`
}
