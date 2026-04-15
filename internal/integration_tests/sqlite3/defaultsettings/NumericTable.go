package dto

import (
	"database/sql"
)

type NumericTable struct {
	N             sql.NullFloat64 `db:"n"`
	NumericNn     float64         `db:"numeric_nn"`
	NumericUnique sql.NullFloat64 `db:"numeric_unique"`
	NumericCheck  sql.NullFloat64 `db:"numeric_check"`
	NumericRef    sql.NullFloat64 `db:"numeric_ref"`
	NumericDef    sql.NullFloat64 `db:"numeric_def"`
	NumericPk     sql.NullFloat64 `db:"numeric_pk"`
}
