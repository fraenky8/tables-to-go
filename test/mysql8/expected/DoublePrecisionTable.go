package dto

import (
	"database/sql"
)

type DoublePrecisionTable struct {
	Col                                sql.NullFloat64 `db:"col"`
	DoublePrecisionNn                  float64         `db:"double_precision_nn"`
	DoublePrecisionNnUnique            float64         `db:"double_precision_nn_unique"`
	DoublePrecisionNnCheck             float64         `db:"double_precision_nn_check"`
	DoublePrecisionNnRef               float64         `db:"double_precision_nn_ref"`
	DoublePrecisionNnDefConst          float64         `db:"double_precision_nn_def_const"`
	DoublePrecisionNnDefFunc           float64         `db:"double_precision_nn_def_func"`
	DoublePrecisionNnUniqueCheck       float64         `db:"double_precision_nn_unique_check"`
	DoublePrecisionUnique              sql.NullFloat64 `db:"double_precision_unique"`
	DoublePrecisionUniqueCheck         sql.NullFloat64 `db:"double_precision_unique_check"`
	DoublePrecisionUniqueRef           sql.NullFloat64 `db:"double_precision_unique_ref"`
	DoublePrecisionUniqueDefConst      sql.NullFloat64 `db:"double_precision_unique_def_const"`
	DoublePrecisionUniqueDefFunc       sql.NullFloat64 `db:"double_precision_unique_def_func"`
	DoublePrecisionCheck               sql.NullFloat64 `db:"double_precision_check"`
	DoublePrecisionCheckRef            sql.NullFloat64 `db:"double_precision_check_ref"`
	DoublePrecisionCheckDefConst       sql.NullFloat64 `db:"double_precision_check_def_const"`
	DoublePrecisionCheckDefFunc        sql.NullFloat64 `db:"double_precision_check_def_func"`
	DoublePrecisionRef                 sql.NullFloat64 `db:"double_precision_ref"`
	DoublePrecisionRefUniqueCheck      sql.NullFloat64 `db:"double_precision_ref_unique_check"`
	DoublePrecisionDefConst            sql.NullFloat64 `db:"double_precision_def_const"`
	DoublePrecisionDefConstUniqueCheck sql.NullFloat64 `db:"double_precision_def_const_unique_check"`
	DoublePrecisionDefFunc             sql.NullFloat64 `db:"double_precision_def_func"`
	DoublePrecisionDefFuncUniqueCheck  sql.NullFloat64 `db:"double_precision_def_func_unique_check"`
}
