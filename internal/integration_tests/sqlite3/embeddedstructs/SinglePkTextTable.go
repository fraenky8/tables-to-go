package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
	"gorm.io/gorm"
)

type SinglePkTextTable struct {
	gorm.Model
	structable.Recorder

	Pk   string         `db:"pk" gorm:"column:pk;primaryKey;not null" stbl:"pk,PRIMARY_KEY"`
	Name sql.NullString `db:"name" gorm:"column:name" stbl:"name"`
}
