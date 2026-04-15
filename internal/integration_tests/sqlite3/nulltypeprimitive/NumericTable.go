package dto

type NumericTable struct {
	N             *float64 `db:"n"`
	NumericNn     float64  `db:"numeric_nn"`
	NumericUnique *float64 `db:"numeric_unique"`
	NumericCheck  *float64 `db:"numeric_check"`
	NumericRef    *float64 `db:"numeric_ref"`
	NumericDef    *float64 `db:"numeric_def"`
	NumericPk     *float64 `db:"numeric_pk"`
}
