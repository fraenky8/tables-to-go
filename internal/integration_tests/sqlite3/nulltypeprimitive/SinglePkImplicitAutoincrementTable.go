package dto

type SinglePkImplicitAutoincrementTable struct {
	Pk   int     `db:"pk"`
	Name *string `db:"name"`
}
