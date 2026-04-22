package dto

import (
	"database/sql"
)

type SinglePkTextTable struct {
	Pk   string         `db:"pk" gorm:"column:pk;primaryKey;not null"`
	Name sql.NullString `db:"name" gorm:"column:name"`
}
