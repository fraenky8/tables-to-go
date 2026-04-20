package dto

type SinglePkExplicitAutoincrementTable struct {
	Pk   int     `db:"pk"`
	Name *string `db:"name"`
}
