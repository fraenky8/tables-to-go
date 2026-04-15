package dto

type IntegerTable struct {
	I               *string `db:"i"`
	IntegerNn       string  `db:"integer_nn"`
	IntegerUnique   *string `db:"integer_unique"`
	IntegerCheck    *string `db:"integer_check"`
	IntegerRef      *string `db:"integer_ref"`
	IntegerDefConst *string `db:"integer_def_const"`
	IntegerPk       *string `db:"integer_pk"`
}
