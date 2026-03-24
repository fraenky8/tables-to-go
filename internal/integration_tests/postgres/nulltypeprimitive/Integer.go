package dto

type Integer struct {
	Integer                    *int `db:"integer"`
	IntegerNn                  int  `db:"integer_nn"`
	IntegerNnUnique            int  `db:"integer_nn_unique"`
	IntegerNnCheck             int  `db:"integer_nn_check"`
	IntegerNnRef               int  `db:"integer_nn_ref"`
	IntegerNnDefConst          int  `db:"integer_nn_def_const"`
	IntegerNnDefFunc           int  `db:"integer_nn_def_func"`
	IntegerNnUniqueCheck       int  `db:"integer_nn_unique_check"`
	IntegerUnique              *int `db:"integer_unique"`
	IntegerUniqueCheck         *int `db:"integer_unique_check"`
	IntegerUniqueRef           *int `db:"integer_unique_ref"`
	IntegerUniqueDefConst      *int `db:"integer_unique_def_const"`
	IntegerUniqueDefFunc       *int `db:"integer_unique_def_func"`
	IntegerCheck               *int `db:"integer_check"`
	IntegerCheckRef            *int `db:"integer_check_ref"`
	IntegerCheckDefConst       *int `db:"integer_check_def_const"`
	IntegerCheckDefFunc        *int `db:"integer_check_def_func"`
	IntegerRef                 *int `db:"integer_ref"`
	IntegerRefDefConst         *int `db:"integer_ref_def_const"`
	IntegerRefDefFunc          *int `db:"integer_ref_def_func"`
	IntegerRefUniqueCheck      *int `db:"integer_ref_unique_check"`
	IntegerDefConst            *int `db:"integer_def_const"`
	IntegerDefConstUniqueCheck *int `db:"integer_def_const_unique_check"`
	IntegerDefFunc             *int `db:"integer_def_func"`
	IntegerDefFuncUniqueCheck  *int `db:"integer_def_func_unique_check"`
}
