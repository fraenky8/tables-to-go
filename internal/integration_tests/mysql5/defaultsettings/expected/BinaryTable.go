package dto

import (
	"database/sql"
)

type BinaryTable struct {
	Col                       sql.NullString `db:"col"`
	BinaryCap                 sql.NullString `db:"binary_cap"`
	BinaryNn                  string         `db:"binary_nn"`
	BinaryNnUnique            string         `db:"binary_nn_unique"`
	BinaryNnCheckCmp          string         `db:"binary_nn_check_cmp"`
	BinaryNnCheckFn           string         `db:"binary_nn_check_fn"`
	BinaryNnRef               string         `db:"binary_nn_ref"`
	BinaryNnDefConst          string         `db:"binary_nn_def_const"`
	BinaryNnDefFunc           string         `db:"binary_nn_def_func"`
	BinaryNnUniqueCheck       string         `db:"binary_nn_unique_check"`
	BinaryUnique              sql.NullString `db:"binary_unique"`
	BinaryUniqueCheck         sql.NullString `db:"binary_unique_check"`
	BinaryUniqueRef           sql.NullString `db:"binary_unique_ref"`
	BinaryUniqueDefConst      sql.NullString `db:"binary_unique_def_const"`
	BinaryUniqueDefFunc       sql.NullString `db:"binary_unique_def_func"`
	BinaryCheck               sql.NullString `db:"binary_check"`
	BinaryCheckRef            sql.NullString `db:"binary_check_ref"`
	BinaryCheckDefConst       sql.NullString `db:"binary_check_def_const"`
	BinaryCheckDefFunc        sql.NullString `db:"binary_check_def_func"`
	BinaryRef                 sql.NullString `db:"binary_ref"`
	BinaryRefUniqueCheck      sql.NullString `db:"binary_ref_unique_check"`
	BinaryDefConst            sql.NullString `db:"binary_def_const"`
	BinaryDefConstUniqueCheck sql.NullString `db:"binary_def_const_unique_check"`
	BinaryDefFunc             sql.NullString `db:"binary_def_func"`
	BinaryDefFuncUniqueCheck  sql.NullString `db:"binary_def_func_unique_check"`
}
