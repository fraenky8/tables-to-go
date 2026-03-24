package dto

type Int2 struct {
	Int2                    *int `db:"int2"`
	Int2Nn                  int  `db:"int2_nn"`
	Int2NnUnique            int  `db:"int2_nn_unique"`
	Int2NnCheck             int  `db:"int2_nn_check"`
	Int2NnRef               int  `db:"int2_nn_ref"`
	Int2NnDefConst          int  `db:"int2_nn_def_const"`
	Int2NnDefFunc           int  `db:"int2_nn_def_func"`
	Int2NnUniqueCheck       int  `db:"int2_nn_unique_check"`
	Int2Unique              *int `db:"int2_unique"`
	Int2UniqueCheck         *int `db:"int2_unique_check"`
	Int2UniqueRef           *int `db:"int2_unique_ref"`
	Int2UniqueDefConst      *int `db:"int2_unique_def_const"`
	Int2UniqueDefFunc       *int `db:"int2_unique_def_func"`
	Int2Check               *int `db:"int2_check"`
	Int2CheckRef            *int `db:"int2_check_ref"`
	Int2CheckDefConst       *int `db:"int2_check_def_const"`
	Int2CheckDefFunc        *int `db:"int2_check_def_func"`
	Int2Ref                 *int `db:"int2_ref"`
	Int2RefDefConst         *int `db:"int2_ref_def_const"`
	Int2RefDefFunc          *int `db:"int2_ref_def_func"`
	Int2RefUniqueCheck      *int `db:"int2_ref_unique_check"`
	Int2DefConst            *int `db:"int2_def_const"`
	Int2DefConstUniqueCheck *int `db:"int2_def_const_unique_check"`
	Int2DefFunc             *int `db:"int2_def_func"`
	Int2DefFuncUniqueCheck  *int `db:"int2_def_func_unique_check"`
}
