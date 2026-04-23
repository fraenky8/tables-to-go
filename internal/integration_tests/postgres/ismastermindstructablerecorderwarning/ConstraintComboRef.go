package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type ConstraintComboRef struct {
	structable.Recorder

	ConstraintComboRef sql.NullFloat64 `db:"constraint_combo_ref"`
}
