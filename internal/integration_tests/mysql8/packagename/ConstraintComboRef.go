package models

import (
	"database/sql"
)

type ConstraintComboRef struct {
	ConstraintComboRef sql.NullFloat64 `db:"constraint_combo_ref"`
}
