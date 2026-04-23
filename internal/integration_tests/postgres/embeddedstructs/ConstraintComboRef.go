package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
	"gorm.io/gorm"
)

type ConstraintComboRef struct {
	gorm.Model
	structable.Recorder

	ConstraintComboRef sql.NullFloat64 `db:"constraint_combo_ref" gorm:"column:constraint_combo_ref" stbl:"constraint_combo_ref"`
}
