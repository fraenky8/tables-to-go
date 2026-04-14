package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type VarcharTable struct {
	Col                        sql.NullString `stbl:"col"`
	VarcharCap                 sql.NullString `stbl:"varchar_cap"`
	VarcharNn                  string         `stbl:"varchar_nn"`
	VarcharNnUnique            string         `stbl:"varchar_nn_unique,PRIMARY_KEY"`
	VarcharNnCheckCmp          string         `stbl:"varchar_nn_check_cmp"`
	VarcharNnCheckFn           string         `stbl:"varchar_nn_check_fn"`
	VarcharNnRef               string         `stbl:"varchar_nn_ref"`
	VarcharNnDefConst          string         `stbl:"varchar_nn_def_const"`
	VarcharNnDefFunc           string         `stbl:"varchar_nn_def_func"`
	VarcharNnUniqueCheck       string         `stbl:"varchar_nn_unique_check"`
	VarcharUnique              sql.NullString `stbl:"varchar_unique"`
	VarcharUniqueCheck         sql.NullString `stbl:"varchar_unique_check"`
	VarcharUniqueRef           sql.NullString `stbl:"varchar_unique_ref"`
	VarcharUniqueDefConst      sql.NullString `stbl:"varchar_unique_def_const"`
	VarcharUniqueDefFunc       sql.NullString `stbl:"varchar_unique_def_func"`
	VarcharCheck               sql.NullString `stbl:"varchar_check"`
	VarcharCheckRef            sql.NullString `stbl:"varchar_check_ref"`
	VarcharCheckDefConst       sql.NullString `stbl:"varchar_check_def_const"`
	VarcharCheckDefFunc        sql.NullString `stbl:"varchar_check_def_func"`
	VarcharRef                 sql.NullString `stbl:"varchar_ref"`
	VarcharRefUniqueCheck      sql.NullString `stbl:"varchar_ref_unique_check"`
	VarcharDefConst            sql.NullString `stbl:"varchar_def_const"`
	VarcharDefConstUniqueCheck sql.NullString `stbl:"varchar_def_const_unique_check"`
	VarcharDefFunc             sql.NullString `stbl:"varchar_def_func"`
	VarcharDefFuncUniqueCheck  sql.NullString `stbl:"varchar_def_func_unique_check"`

	structable.Recorder
}
