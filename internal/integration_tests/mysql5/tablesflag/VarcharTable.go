package dto

import (
	"database/sql"
)

type VarcharTable struct {
	Col                        sql.NullString `db:"col"`
	VarcharCap                 sql.NullString `db:"varchar_cap"`
	VarcharNn                  string         `db:"varchar_nn"`
	VarcharNnUnique            string         `db:"varchar_nn_unique"`
	VarcharNnCheckCmp          string         `db:"varchar_nn_check_cmp"`
	VarcharNnCheckFn           string         `db:"varchar_nn_check_fn"`
	VarcharNnRef               string         `db:"varchar_nn_ref"`
	VarcharNnDefConst          string         `db:"varchar_nn_def_const"`
	VarcharNnDefFunc           string         `db:"varchar_nn_def_func"`
	VarcharNnUniqueCheck       string         `db:"varchar_nn_unique_check"`
	VarcharUnique              sql.NullString `db:"varchar_unique"`
	VarcharUniqueCheck         sql.NullString `db:"varchar_unique_check"`
	VarcharUniqueRef           sql.NullString `db:"varchar_unique_ref"`
	VarcharUniqueDefConst      sql.NullString `db:"varchar_unique_def_const"`
	VarcharUniqueDefFunc       sql.NullString `db:"varchar_unique_def_func"`
	VarcharCheck               sql.NullString `db:"varchar_check"`
	VarcharCheckRef            sql.NullString `db:"varchar_check_ref"`
	VarcharCheckDefConst       sql.NullString `db:"varchar_check_def_const"`
	VarcharCheckDefFunc        sql.NullString `db:"varchar_check_def_func"`
	VarcharRef                 sql.NullString `db:"varchar_ref"`
	VarcharRefUniqueCheck      sql.NullString `db:"varchar_ref_unique_check"`
	VarcharDefConst            sql.NullString `db:"varchar_def_const"`
	VarcharDefConstUniqueCheck sql.NullString `db:"varchar_def_const_unique_check"`
	VarcharDefFunc             sql.NullString `db:"varchar_def_func"`
	VarcharDefFuncUniqueCheck  sql.NullString `db:"varchar_def_func_unique_check"`
}
