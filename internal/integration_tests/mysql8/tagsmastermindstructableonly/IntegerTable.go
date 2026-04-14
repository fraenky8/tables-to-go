package dto

import (
	"database/sql"
)

type IntegerTable struct {
	I                          sql.NullInt64 `stbl:"i"`
	IntegerNn                  int           `stbl:"integer_nn"`
	IntegerNnUnique            int           `stbl:"integer_nn_unique,PRIMARY_KEY"`
	IntegerNnCheck             int           `stbl:"integer_nn_check"`
	IntegerUnique              sql.NullInt64 `stbl:"integer_unique"`
	IntegerUniqueCheck         sql.NullInt64 `stbl:"integer_unique_check"`
	IntegerUniqueRef           sql.NullInt64 `stbl:"integer_unique_ref"`
	IntegerUniqueDefConst      sql.NullInt64 `stbl:"integer_unique_def_const"`
	IntegerUniqueDefFunc       sql.NullInt64 `stbl:"integer_unique_def_func"`
	IntegerCheck               sql.NullInt64 `stbl:"integer_check"`
	IntegerCheckRef            sql.NullInt64 `stbl:"integer_check_ref"`
	IntegerCheckDefConst       sql.NullInt64 `stbl:"integer_check_def_const"`
	IntegerCheckDefFunc        sql.NullInt64 `stbl:"integer_check_def_func"`
	IntegerRef                 sql.NullInt64 `stbl:"integer_ref"`
	IntegerRefUniqueCheck      sql.NullInt64 `stbl:"integer_ref_unique_check"`
	IntegerDefConst            sql.NullInt64 `stbl:"integer_def_const"`
	IntegerDefConstUniqueCheck sql.NullInt64 `stbl:"integer_def_const_unique_check"`
	IntegerDefFunc             sql.NullInt64 `stbl:"integer_def_func"`
	IntegerDefFuncUniqueCheck  sql.NullInt64 `stbl:"integer_def_func_unique_check"`
}
