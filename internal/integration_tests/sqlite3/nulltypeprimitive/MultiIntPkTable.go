package dto

type MultiIntPkTable struct {
	PkA  int     `db:"pk_a"`
	PkB  int     `db:"pk_b"`
	Name *string `db:"name"`
}
