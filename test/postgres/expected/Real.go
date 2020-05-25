package dto

import (
	"database/sql"
)

type Real struct {
	Real                    sql.NullFloat64 `db:"real"`
	RealNn                  float64         `db:"real_nn"`
	RealNnUnique            float64         `db:"real_nn_unique"`
	RealNnCheck             float64         `db:"real_nn_check"`
	RealNnRef               float64         `db:"real_nn_ref"`
	RealNnDefConst          float64         `db:"real_nn_def_const"`
	RealNnDefFunc           float64         `db:"real_nn_def_func"`
	RealNnUniqueCheck       float64         `db:"real_nn_unique_check"`
	RealUnique              sql.NullFloat64 `db:"real_unique"`
	RealUniqueCheck         sql.NullFloat64 `db:"real_unique_check"`
	RealUniqueRef           sql.NullFloat64 `db:"real_unique_ref"`
	RealUniqueDefConst      sql.NullFloat64 `db:"real_unique_def_const"`
	RealUniqueDefFunc       sql.NullFloat64 `db:"real_unique_def_func"`
	RealCheck               sql.NullFloat64 `db:"real_check"`
	RealCheckRef            sql.NullFloat64 `db:"real_check_ref"`
	RealCheckDefConst       sql.NullFloat64 `db:"real_check_def_const"`
	RealCheckDefFunc        sql.NullFloat64 `db:"real_check_def_func"`
	RealRef                 sql.NullFloat64 `db:"real_ref"`
	RealRefDefConst         sql.NullFloat64 `db:"real_ref_def_const"`
	RealRefDefFunc          sql.NullFloat64 `db:"real_ref_def_func"`
	RealRefUniqueCheck      sql.NullFloat64 `db:"real_ref_unique_check"`
	RealDefConst            sql.NullFloat64 `db:"real_def_const"`
	RealDefConstUniqueCheck sql.NullFloat64 `db:"real_def_const_unique_check"`
	RealDefFunc             sql.NullFloat64 `db:"real_def_func"`
	RealDefFuncUniqueCheck  sql.NullFloat64 `db:"real_def_func_unique_check"`
}
