package dto

type LongtextTable struct {
	Col                   *string `db:"col"`
	LongtextDefConst      *string `db:"longtext_def_const"`
	LongtextDefFunc       *string `db:"longtext_def_func"`
	LongtextRef           *string `db:"longtext_ref"`
	LongtextNn            string  `db:"longtext_nn"`
	LongtextNnCheckCmp    string  `db:"longtext_nn_check_cmp"`
	LongtextNnCheckFn     string  `db:"longtext_nn_check_fn"`
	LongtextNnRef         string  `db:"longtext_nn_ref"`
	LongtextNnDefConst    string  `db:"longtext_nn_def_const"`
	LongtextNnDefFunc     string  `db:"longtext_nn_def_func"`
	LongtextCheck         *string `db:"longtext_check"`
	LongtextCheckRef      *string `db:"longtext_check_ref"`
	LongtextCheckDefConst *string `db:"longtext_check_def_const"`
	LongtextCheckDefFunc  *string `db:"longtext_check_def_func"`
}
