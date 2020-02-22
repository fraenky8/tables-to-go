package dto

import (
	"database/sql"
)

type Int8 struct {
	Int8                    sql.NullInt64 `db:"int8"`
	Int8Nn                  int           `db:"int8_nn"`
	Int8NnUnique            int           `db:"int8_nn_unique"`
	Int8NnCheck             int           `db:"int8_nn_check"`
	Int8NnRef               int           `db:"int8_nn_ref"`
	Int8NnDefConst          int           `db:"int8_nn_def_const"`
	Int8NnDefFunc           int           `db:"int8_nn_def_func"`
	Int8NnUniqueCheck       int           `db:"int8_nn_unique_check"`
	Int8Unique              sql.NullInt64 `db:"int8_unique"`
	Int8UniqueCheck         sql.NullInt64 `db:"int8_unique_check"`
	Int8UniqueRef           sql.NullInt64 `db:"int8_unique_ref"`
	Int8UniqueDefConst      sql.NullInt64 `db:"int8_unique_def_const"`
	Int8UniqueDefFunc       sql.NullInt64 `db:"int8_unique_def_func"`
	Int8Check               sql.NullInt64 `db:"int8_check"`
	Int8CheckRef            sql.NullInt64 `db:"int8_check_ref"`
	Int8CheckDefConst       sql.NullInt64 `db:"int8_check_def_const"`
	Int8CheckDefFunc        sql.NullInt64 `db:"int8_check_def_func"`
	Int8Ref                 sql.NullInt64 `db:"int8_ref"`
	Int8RefDefConst         sql.NullInt64 `db:"int8_ref_def_const"`
	Int8RefDefFunc          sql.NullInt64 `db:"int8_ref_def_func"`
	Int8RefUniqueCheck      sql.NullInt64 `db:"int8_ref_unique_check"`
	Int8DefConst            sql.NullInt64 `db:"int8_def_const"`
	Int8DefConstUniqueCheck sql.NullInt64 `db:"int8_def_const_unique_check"`
	Int8DefFunc             sql.NullInt64 `db:"int8_def_func"`
	Int8DefFuncUniqueCheck  sql.NullInt64 `db:"int8_def_func_unique_check"`
}
