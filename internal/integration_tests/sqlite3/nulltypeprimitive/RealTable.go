package dto

type RealTable struct {
	R            *string `db:"r"`
	RealNn       string  `db:"real_nn"`
	RealUnique   *string `db:"real_unique"`
	RealCheck    *string `db:"real_check"`
	RealRef      *string `db:"real_ref"`
	RealDefConst *string `db:"real_def_const"`
	RealPk       *string `db:"real_pk"`
}
