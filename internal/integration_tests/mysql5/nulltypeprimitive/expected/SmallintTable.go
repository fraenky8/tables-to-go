package dto

type SmallintTable struct {
	I                           *int `db:"i"`
	SmallintNn                  int  `db:"smallint_nn"`
	SmallintNnUnique            int  `db:"smallint_nn_unique"`
	SmallintNnCheck             int  `db:"smallint_nn_check"`
	SmallintUnique              *int `db:"smallint_unique"`
	SmallintUniqueCheck         *int `db:"smallint_unique_check"`
	SmallintUniqueRef           *int `db:"smallint_unique_ref"`
	SmallintUniqueDefConst      *int `db:"smallint_unique_def_const"`
	SmallintUniqueDefFunc       *int `db:"smallint_unique_def_func"`
	SmallintCheck               *int `db:"smallint_check"`
	SmallintCheckRef            *int `db:"smallint_check_ref"`
	SmallintCheckDefConst       *int `db:"smallint_check_def_const"`
	SmallintCheckDefFunc        *int `db:"smallint_check_def_func"`
	SmallintRef                 *int `db:"smallint_ref"`
	SmallintRefUniqueCheck      *int `db:"smallint_ref_unique_check"`
	SmallintDefConst            *int `db:"smallint_def_const"`
	SmallintDefConstUniqueCheck *int `db:"smallint_def_const_unique_check"`
	SmallintDefFunc             *int `db:"smallint_def_func"`
	SmallintDefFuncUniqueCheck  *int `db:"smallint_def_func_unique_check"`
}
