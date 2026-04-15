package dto

import (
	"database/sql"
)

type RealTable struct {
	R            sql.NullFloat64 `db:"r"`
	RealNn       float64         `db:"real_nn"`
	RealUnique   sql.NullFloat64 `db:"real_unique"`
	RealCheck    sql.NullFloat64 `db:"real_check"`
	RealRef      sql.NullFloat64 `db:"real_ref"`
	RealDefConst sql.NullFloat64 `db:"real_def_const"`
	RealPk       sql.NullFloat64 `db:"real_pk"`
}
