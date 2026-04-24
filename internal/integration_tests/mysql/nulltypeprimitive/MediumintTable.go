package dto

type MediumintTable struct {
	I                            *int `db:"i"`
	MediumintNn                  int  `db:"mediumint_nn"`
	MediumintNnUnique            int  `db:"mediumint_nn_unique"`
	MediumintNnCheck             int  `db:"mediumint_nn_check"`
	MediumintUnique              *int `db:"mediumint_unique"`
	MediumintUniqueCheck         *int `db:"mediumint_unique_check"`
	MediumintUniqueRef           *int `db:"mediumint_unique_ref"`
	MediumintUniqueDefConst      *int `db:"mediumint_unique_def_const"`
	MediumintUniqueDefFunc       *int `db:"mediumint_unique_def_func"`
	MediumintCheck               *int `db:"mediumint_check"`
	MediumintCheckRef            *int `db:"mediumint_check_ref"`
	MediumintCheckDefConst       *int `db:"mediumint_check_def_const"`
	MediumintCheckDefFunc        *int `db:"mediumint_check_def_func"`
	MediumintRef                 *int `db:"mediumint_ref"`
	MediumintRefUniqueCheck      *int `db:"mediumint_ref_unique_check"`
	MediumintDefConst            *int `db:"mediumint_def_const"`
	MediumintDefConstUniqueCheck *int `db:"mediumint_def_const_unique_check"`
	MediumintDefFunc             *int `db:"mediumint_def_func"`
	MediumintDefFuncUniqueCheck  *int `db:"mediumint_def_func_unique_check"`
}
