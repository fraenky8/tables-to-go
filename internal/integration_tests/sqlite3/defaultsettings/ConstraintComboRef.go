package dto

import (
	"database/sql"
)

type ConstraintComboRef struct {
	ConstraintComboRef sql.NullString `db:"constraint_combo_ref"`
}
