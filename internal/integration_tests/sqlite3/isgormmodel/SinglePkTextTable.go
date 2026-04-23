package dto

import (
	"database/sql"

	"gorm.io/gorm"
)

type SinglePkTextTable struct {
	gorm.Model

	Pk   string         `db:"pk" gorm:"column:pk;primaryKey;not null"`
	Name sql.NullString `db:"name" gorm:"column:name"`
}
