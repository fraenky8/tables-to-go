package dto

type Int8 struct {
	Int8                    *int `db:"int8"`
	Int8Nn                  int  `db:"int8_nn"`
	Int8NnUnique            int  `db:"int8_nn_unique"`
	Int8NnCheck             int  `db:"int8_nn_check"`
	Int8NnRef               int  `db:"int8_nn_ref"`
	Int8NnDefConst          int  `db:"int8_nn_def_const"`
	Int8NnDefFunc           int  `db:"int8_nn_def_func"`
	Int8NnUniqueCheck       int  `db:"int8_nn_unique_check"`
	Int8Unique              *int `db:"int8_unique"`
	Int8UniqueCheck         *int `db:"int8_unique_check"`
	Int8UniqueRef           *int `db:"int8_unique_ref"`
	Int8UniqueDefConst      *int `db:"int8_unique_def_const"`
	Int8UniqueDefFunc       *int `db:"int8_unique_def_func"`
	Int8Check               *int `db:"int8_check"`
	Int8CheckRef            *int `db:"int8_check_ref"`
	Int8CheckDefConst       *int `db:"int8_check_def_const"`
	Int8CheckDefFunc        *int `db:"int8_check_def_func"`
	Int8Ref                 *int `db:"int8_ref"`
	Int8RefDefConst         *int `db:"int8_ref_def_const"`
	Int8RefDefFunc          *int `db:"int8_ref_def_func"`
	Int8RefUniqueCheck      *int `db:"int8_ref_unique_check"`
	Int8DefConst            *int `db:"int8_def_const"`
	Int8DefConstUniqueCheck *int `db:"int8_def_const_unique_check"`
	Int8DefFunc             *int `db:"int8_def_func"`
	Int8DefFuncUniqueCheck  *int `db:"int8_def_func_unique_check"`
}
