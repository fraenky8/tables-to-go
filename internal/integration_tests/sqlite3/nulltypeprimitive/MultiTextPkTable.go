package dto

type MultiTextPkTable struct {
	PkA  string  `db:"pk_a"`
	PkB  string  `db:"pk_b"`
	Name *string `db:"name"`
}
