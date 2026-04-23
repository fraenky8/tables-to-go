package dto

import (
	"database/sql"

	"gorm.io/gorm"
)

type ConstraintComboRef struct {
	gorm.Model

	ConstraintComboRef sql.NullFloat64 `db:"constraint_combo_ref"`
}
