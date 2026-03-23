package dto

type TinyintTable struct {
	I                          *int `db:"i"`
	TinyintNn                  int  `db:"tinyint_nn"`
	TinyintNnUnique            int  `db:"tinyint_nn_unique"`
	TinyintNnCheck             int  `db:"tinyint_nn_check"`
	TinyintUnique              *int `db:"tinyint_unique"`
	TinyintUniqueCheck         *int `db:"tinyint_unique_check"`
	TinyintUniqueRef           *int `db:"tinyint_unique_ref"`
	TinyintUniqueDefConst      *int `db:"tinyint_unique_def_const"`
	TinyintUniqueDefFunc       *int `db:"tinyint_unique_def_func"`
	TinyintCheck               *int `db:"tinyint_check"`
	TinyintCheckRef            *int `db:"tinyint_check_ref"`
	TinyintCheckDefConst       *int `db:"tinyint_check_def_const"`
	TinyintCheckDefFunc        *int `db:"tinyint_check_def_func"`
	TinyintRef                 *int `db:"tinyint_ref"`
	TinyintRefUniqueCheck      *int `db:"tinyint_ref_unique_check"`
	TinyintDefConst            *int `db:"tinyint_def_const"`
	TinyintDefConstUniqueCheck *int `db:"tinyint_def_const_unique_check"`
	TinyintDefFunc             *int `db:"tinyint_def_func"`
	TinyintDefFuncUniqueCheck  *int `db:"tinyint_def_func_unique_check"`
}
