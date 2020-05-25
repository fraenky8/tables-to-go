package dto

import (
	"database/sql"
)

type Int2 struct {
	Int2                    sql.NullInt64 `db:"int2"`
	Int2Nn                  int           `db:"int2_nn"`
	Int2NnUnique            int           `db:"int2_nn_unique"`
	Int2NnCheck             int           `db:"int2_nn_check"`
	Int2NnRef               int           `db:"int2_nn_ref"`
	Int2NnDefConst          int           `db:"int2_nn_def_const"`
	Int2NnDefFunc           int           `db:"int2_nn_def_func"`
	Int2NnUniqueCheck       int           `db:"int2_nn_unique_check"`
	Int2Unique              sql.NullInt64 `db:"int2_unique"`
	Int2UniqueCheck         sql.NullInt64 `db:"int2_unique_check"`
	Int2UniqueRef           sql.NullInt64 `db:"int2_unique_ref"`
	Int2UniqueDefConst      sql.NullInt64 `db:"int2_unique_def_const"`
	Int2UniqueDefFunc       sql.NullInt64 `db:"int2_unique_def_func"`
	Int2Check               sql.NullInt64 `db:"int2_check"`
	Int2CheckRef            sql.NullInt64 `db:"int2_check_ref"`
	Int2CheckDefConst       sql.NullInt64 `db:"int2_check_def_const"`
	Int2CheckDefFunc        sql.NullInt64 `db:"int2_check_def_func"`
	Int2Ref                 sql.NullInt64 `db:"int2_ref"`
	Int2RefDefConst         sql.NullInt64 `db:"int2_ref_def_const"`
	Int2RefDefFunc          sql.NullInt64 `db:"int2_ref_def_func"`
	Int2RefUniqueCheck      sql.NullInt64 `db:"int2_ref_unique_check"`
	Int2DefConst            sql.NullInt64 `db:"int2_def_const"`
	Int2DefConstUniqueCheck sql.NullInt64 `db:"int2_def_const_unique_check"`
	Int2DefFunc             sql.NullInt64 `db:"int2_def_func"`
	Int2DefFuncUniqueCheck  sql.NullInt64 `db:"int2_def_func_unique_check"`
}
