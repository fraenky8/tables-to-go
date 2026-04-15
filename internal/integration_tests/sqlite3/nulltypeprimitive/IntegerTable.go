package dto

type IntegerTable struct {
	I               *int `db:"i"`
	IntegerNn       int  `db:"integer_nn"`
	IntegerUnique   *int `db:"integer_unique"`
	IntegerCheck    *int `db:"integer_check"`
	IntegerRef      *int `db:"integer_ref"`
	IntegerDefConst *int `db:"integer_def_const"`
	IntegerPk       *int `db:"integer_pk"`
}
