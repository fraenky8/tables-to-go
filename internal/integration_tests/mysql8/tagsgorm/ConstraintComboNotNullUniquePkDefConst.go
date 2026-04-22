package dto

type ConstraintComboNotNullUniquePkDefConst struct {
	ConstraintComboNotNullUniquePkDefConst float64 `db:"constraint_combo_not_null_unique_pk_def_const" gorm:"column:constraint_combo_not_null_unique_pk_def_const;primaryKey;not null;default:42"`
}
