package dto

import (
	"database/sql"
)

type MediumintTable struct {
	I                            sql.NullInt64 `db:"i"`
	MediumintNn                  int           `db:"mediumint_nn"`
	MediumintNnUnique            int           `db:"mediumint_nn_unique"`
	MediumintNnCheck             int           `db:"mediumint_nn_check"`
	MediumintUnique              sql.NullInt64 `db:"mediumint_unique"`
	MediumintUniqueCheck         sql.NullInt64 `db:"mediumint_unique_check"`
	MediumintUniqueRef           sql.NullInt64 `db:"mediumint_unique_ref"`
	MediumintUniqueDefConst      sql.NullInt64 `db:"mediumint_unique_def_const"`
	MediumintUniqueDefFunc       sql.NullInt64 `db:"mediumint_unique_def_func"`
	MediumintCheck               sql.NullInt64 `db:"mediumint_check"`
	MediumintCheckRef            sql.NullInt64 `db:"mediumint_check_ref"`
	MediumintCheckDefConst       sql.NullInt64 `db:"mediumint_check_def_const"`
	MediumintCheckDefFunc        sql.NullInt64 `db:"mediumint_check_def_func"`
	MediumintRef                 sql.NullInt64 `db:"mediumint_ref"`
	MediumintRefUniqueCheck      sql.NullInt64 `db:"mediumint_ref_unique_check"`
	MediumintDefConst            sql.NullInt64 `db:"mediumint_def_const"`
	MediumintDefConstUniqueCheck sql.NullInt64 `db:"mediumint_def_const_unique_check"`
	MediumintDefFunc             sql.NullInt64 `db:"mediumint_def_func"`
	MediumintDefFuncUniqueCheck  sql.NullInt64 `db:"mediumint_def_func_unique_check"`
}
