package dto

import (
	"database/sql"
)

type SinglePkImplicitAutoincrementTable struct {
	Pk   int            `db:"pk" gorm:"column:pk;primaryKey;autoIncrement;not null"`
	Name sql.NullString `db:"name" gorm:"column:name"`
}
