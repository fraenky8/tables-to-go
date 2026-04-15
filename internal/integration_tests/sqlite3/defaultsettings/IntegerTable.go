package dto

import (
	"database/sql"
)

type IntegerTable struct {
	I               sql.NullInt64 `db:"i"`
	IntegerNn       int           `db:"integer_nn"`
	IntegerUnique   sql.NullInt64 `db:"integer_unique"`
	IntegerCheck    sql.NullInt64 `db:"integer_check"`
	IntegerRef      sql.NullInt64 `db:"integer_ref"`
	IntegerDefConst sql.NullInt64 `db:"integer_def_const"`
	IntegerPk       sql.NullInt64 `db:"integer_pk"`
}
