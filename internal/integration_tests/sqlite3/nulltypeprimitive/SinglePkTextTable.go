package dto

type SinglePkTextTable struct {
	Pk   string  `db:"pk"`
	Name *string `db:"name"`
}
