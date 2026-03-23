package dto

type Int4 struct {
	Int4                    *int `db:"int4"`
	Int4Nn                  int  `db:"int4_nn"`
	Int4NnUnique            int  `db:"int4_nn_unique"`
	Int4NnCheck             int  `db:"int4_nn_check"`
	Int4NnRef               int  `db:"int4_nn_ref"`
	Int4NnDefConst          int  `db:"int4_nn_def_const"`
	Int4NnDefFunc           int  `db:"int4_nn_def_func"`
	Int4NnUniqueCheck       int  `db:"int4_nn_unique_check"`
	Int4Unique              *int `db:"int4_unique"`
	Int4UniqueCheck         *int `db:"int4_unique_check"`
	Int4UniqueRef           *int `db:"int4_unique_ref"`
	Int4UniqueDefConst      *int `db:"int4_unique_def_const"`
	Int4UniqueDefFunc       *int `db:"int4_unique_def_func"`
	Int4Check               *int `db:"int4_check"`
	Int4CheckRef            *int `db:"int4_check_ref"`
	Int4CheckDefConst       *int `db:"int4_check_def_const"`
	Int4CheckDefFunc        *int `db:"int4_check_def_func"`
	Int4Ref                 *int `db:"int4_ref"`
	Int4RefDefConst         *int `db:"int4_ref_def_const"`
	Int4RefDefFunc          *int `db:"int4_ref_def_func"`
	Int4RefUniqueCheck      *int `db:"int4_ref_unique_check"`
	Int4DefConst            *int `db:"int4_def_const"`
	Int4DefConstUniqueCheck *int `db:"int4_def_const_unique_check"`
	Int4DefFunc             *int `db:"int4_def_func"`
	Int4DefFuncUniqueCheck  *int `db:"int4_def_func_unique_check"`
}
