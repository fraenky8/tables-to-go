package dto

import (
	"database/sql"
)

type BigintTable struct {
	I                         sql.NullInt64 `db:"i"`
	BigintNn                  int           `db:"bigint_nn"`
	BigintNnUnique            int           `db:"bigint_nn_unique"`
	BigintNnCheck             int           `db:"bigint_nn_check"`
	BigintUnique              sql.NullInt64 `db:"bigint_unique"`
	BigintUniqueCheck         sql.NullInt64 `db:"bigint_unique_check"`
	BigintUniqueRef           sql.NullInt64 `db:"bigint_unique_ref"`
	BigintUniqueDefConst      sql.NullInt64 `db:"bigint_unique_def_const"`
	BigintUniqueDefFunc       sql.NullInt64 `db:"bigint_unique_def_func"`
	BigintCheck               sql.NullInt64 `db:"bigint_check"`
	BigintCheckRef            sql.NullInt64 `db:"bigint_check_ref"`
	BigintCheckDefConst       sql.NullInt64 `db:"bigint_check_def_const"`
	BigintCheckDefFunc        sql.NullInt64 `db:"bigint_check_def_func"`
	BigintRef                 sql.NullInt64 `db:"bigint_ref"`
	BigintRefUniqueCheck      sql.NullInt64 `db:"bigint_ref_unique_check"`
	BigintDefConst            sql.NullInt64 `db:"bigint_def_const"`
	BigintDefConstUniqueCheck sql.NullInt64 `db:"bigint_def_const_unique_check"`
	BigintDefFunc             sql.NullInt64 `db:"bigint_def_func"`
	BigintDefFuncUniqueCheck  sql.NullInt64 `db:"bigint_def_func_unique_check"`
}
