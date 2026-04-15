package dto

type RealTable struct {
	R            *float64 `db:"r"`
	RealNn       float64  `db:"real_nn"`
	RealUnique   *float64 `db:"real_unique"`
	RealCheck    *float64 `db:"real_check"`
	RealRef      *float64 `db:"real_ref"`
	RealDefConst *float64 `db:"real_def_const"`
	RealPk       *float64 `db:"real_pk"`
}
