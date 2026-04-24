package dto

type MediumblobTable struct {
	Col                     *string `db:"col"`
	MediumblobDefConst      *string `db:"mediumblob_def_const"`
	MediumblobDefFunc       *string `db:"mediumblob_def_func"`
	MediumblobRef           *string `db:"mediumblob_ref"`
	MediumblobNn            string  `db:"mediumblob_nn"`
	MediumblobNnCheckCmp    string  `db:"mediumblob_nn_check_cmp"`
	MediumblobNnCheckFn     string  `db:"mediumblob_nn_check_fn"`
	MediumblobNnRef         string  `db:"mediumblob_nn_ref"`
	MediumblobNnDefConst    string  `db:"mediumblob_nn_def_const"`
	MediumblobNnDefFunc     string  `db:"mediumblob_nn_def_func"`
	MediumblobCheck         *string `db:"mediumblob_check"`
	MediumblobCheckRef      *string `db:"mediumblob_check_ref"`
	MediumblobCheckDefConst *string `db:"mediumblob_check_def_const"`
	MediumblobCheckDefFunc  *string `db:"mediumblob_check_def_func"`
}
