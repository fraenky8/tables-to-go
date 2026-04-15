package dto

import (
	"database/sql"
)

type RealTable struct {
	R            sql.NullString `db:"r"`
	RealNn       string         `db:"real_nn"`
	RealUnique   sql.NullString `db:"real_unique"`
	RealCheck    sql.NullString `db:"real_check"`
	RealRef      sql.NullString `db:"real_ref"`
	RealDefConst sql.NullString `db:"real_def_const"`
	RealPk       sql.NullString `db:"real_pk"`
}
