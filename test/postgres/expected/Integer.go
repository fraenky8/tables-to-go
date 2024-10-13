package dto

import (
	"database/sql"
)

type Integer struct {
	Integer                    sql.NullInt64 `db:"integer"`
	IntegerNn                  int           `db:"integer_nn"`
	IntegerNnUnique            int           `db:"integer_nn_unique"`
	IntegerNnCheck             int           `db:"integer_nn_check"`
	IntegerNnRef               int           `db:"integer_nn_ref"`
	IntegerNnDefConst          int           `db:"integer_nn_def_const"`
	IntegerNnDefFunc           int           `db:"integer_nn_def_func"`
	IntegerNnUniqueCheck       int           `db:"integer_nn_unique_check"`
	IntegerUnique              sql.NullInt64 `db:"integer_unique"`
	IntegerUniqueCheck         sql.NullInt64 `db:"integer_unique_check"`
	IntegerUniqueRef           sql.NullInt64 `db:"integer_unique_ref"`
	IntegerUniqueDefConst      sql.NullInt64 `db:"integer_unique_def_const"`
	IntegerUniqueDefFunc       sql.NullInt64 `db:"integer_unique_def_func"`
	IntegerCheck               sql.NullInt64 `db:"integer_check"`
	IntegerCheckRef            sql.NullInt64 `db:"integer_check_ref"`
	IntegerCheckDefConst       sql.NullInt64 `db:"integer_check_def_const"`
	IntegerCheckDefFunc        sql.NullInt64 `db:"integer_check_def_func"`
	IntegerRef                 sql.NullInt64 `db:"integer_ref"`
	IntegerRefDefConst         sql.NullInt64 `db:"integer_ref_def_const"`
	IntegerRefDefFunc          sql.NullInt64 `db:"integer_ref_def_func"`
	IntegerRefUniqueCheck      sql.NullInt64 `db:"integer_ref_unique_check"`
	IntegerDefConst            sql.NullInt64 `db:"integer_def_const"`
	IntegerDefConstUniqueCheck sql.NullInt64 `db:"integer_def_const_unique_check"`
	IntegerDefFunc             sql.NullInt64 `db:"integer_def_func"`
	IntegerDefFuncUniqueCheck  sql.NullInt64 `db:"integer_def_func_unique_check"`
}
