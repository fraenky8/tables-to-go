package dto

import (
	"database/sql"
)

type Int4 struct {
	Int4                    sql.NullInt64 `db:"int4"`
	Int4Nn                  int           `db:"int4_nn"`
	Int4NnUnique            int           `db:"int4_nn_unique"`
	Int4NnCheck             int           `db:"int4_nn_check"`
	Int4NnRef               int           `db:"int4_nn_ref"`
	Int4NnDefConst          int           `db:"int4_nn_def_const"`
	Int4NnDefFunc           int           `db:"int4_nn_def_func"`
	Int4NnUniqueCheck       int           `db:"int4_nn_unique_check"`
	Int4Unique              sql.NullInt64 `db:"int4_unique"`
	Int4UniqueCheck         sql.NullInt64 `db:"int4_unique_check"`
	Int4UniqueRef           sql.NullInt64 `db:"int4_unique_ref"`
	Int4UniqueDefConst      sql.NullInt64 `db:"int4_unique_def_const"`
	Int4UniqueDefFunc       sql.NullInt64 `db:"int4_unique_def_func"`
	Int4Check               sql.NullInt64 `db:"int4_check"`
	Int4CheckRef            sql.NullInt64 `db:"int4_check_ref"`
	Int4CheckDefConst       sql.NullInt64 `db:"int4_check_def_const"`
	Int4CheckDefFunc        sql.NullInt64 `db:"int4_check_def_func"`
	Int4Ref                 sql.NullInt64 `db:"int4_ref"`
	Int4RefDefConst         sql.NullInt64 `db:"int4_ref_def_const"`
	Int4RefDefFunc          sql.NullInt64 `db:"int4_ref_def_func"`
	Int4RefUniqueCheck      sql.NullInt64 `db:"int4_ref_unique_check"`
	Int4DefConst            sql.NullInt64 `db:"int4_def_const"`
	Int4DefConstUniqueCheck sql.NullInt64 `db:"int4_def_const_unique_check"`
	Int4DefFunc             sql.NullInt64 `db:"int4_def_func"`
	Int4DefFuncUniqueCheck  sql.NullInt64 `db:"int4_def_func_unique_check"`
}
