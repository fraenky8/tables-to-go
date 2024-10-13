package dto

import (
	"database/sql"
)

type Float4 struct {
	Float4                    sql.NullFloat64 `db:"float4"`
	Float4Nn                  float64         `db:"float4_nn"`
	Float4NnUnique            float64         `db:"float4_nn_unique"`
	Float4NnCheck             float64         `db:"float4_nn_check"`
	Float4NnRef               float64         `db:"float4_nn_ref"`
	Float4NnDefConst          float64         `db:"float4_nn_def_const"`
	Float4NnDefFunc           float64         `db:"float4_nn_def_func"`
	Float4NnUniqueCheck       float64         `db:"float4_nn_unique_check"`
	Float4Unique              sql.NullFloat64 `db:"float4_unique"`
	Float4UniqueCheck         sql.NullFloat64 `db:"float4_unique_check"`
	Float4UniqueRef           sql.NullFloat64 `db:"float4_unique_ref"`
	Float4UniqueDefConst      sql.NullFloat64 `db:"float4_unique_def_const"`
	Float4UniqueDefFunc       sql.NullFloat64 `db:"float4_unique_def_func"`
	Float4Check               sql.NullFloat64 `db:"float4_check"`
	Float4CheckRef            sql.NullFloat64 `db:"float4_check_ref"`
	Float4CheckDefConst       sql.NullFloat64 `db:"float4_check_def_const"`
	Float4CheckDefFunc        sql.NullFloat64 `db:"float4_check_def_func"`
	Float4Ref                 sql.NullFloat64 `db:"float4_ref"`
	Float4RefDefConst         sql.NullFloat64 `db:"float4_ref_def_const"`
	Float4RefDefFunc          sql.NullFloat64 `db:"float4_ref_def_func"`
	Float4RefUniqueCheck      sql.NullFloat64 `db:"float4_ref_unique_check"`
	Float4DefConst            sql.NullFloat64 `db:"float4_def_const"`
	Float4DefConstUniqueCheck sql.NullFloat64 `db:"float4_def_const_unique_check"`
	Float4DefFunc             sql.NullFloat64 `db:"float4_def_func"`
	Float4DefFuncUniqueCheck  sql.NullFloat64 `db:"float4_def_func_unique_check"`
}
