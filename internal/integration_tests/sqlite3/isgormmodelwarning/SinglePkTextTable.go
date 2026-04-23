package dto

import (
	"database/sql"

	"gorm.io/gorm"
)

type SinglePkTextTable struct {
	gorm.Model

	Pk   string         `db:"pk"`
	Name sql.NullString `db:"name"`
}
