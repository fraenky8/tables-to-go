package dto

import (
	"database/sql"
)

type IntegerTable struct {
	I               sql.NullString `db:"i"`
	IntegerNn       string         `db:"integer_nn"`
	IntegerUnique   sql.NullString `db:"integer_unique"`
	IntegerCheck    sql.NullString `db:"integer_check"`
	IntegerRef      sql.NullString `db:"integer_ref"`
	IntegerDefConst sql.NullString `db:"integer_def_const"`
	IntegerPk       sql.NullString `db:"integer_pk"`
}
