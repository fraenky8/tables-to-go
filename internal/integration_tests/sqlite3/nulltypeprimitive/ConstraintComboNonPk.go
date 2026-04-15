package dto

type ConstraintComboNonPk struct {
	NotNullUniqueCheckDefConst float64 `db:"not_null_unique_check_def_const"`
	NotNullUniqueRefDefConst   float64 `db:"not_null_unique_ref_def_const"`
	NotNullCheckRefDefConst    float64 `db:"not_null_check_ref_def_const"`
	NotNullUniqueDefConst      float64 `db:"not_null_unique_def_const"`
	NotNullCheckDefConst       float64 `db:"not_null_check_def_const"`
	NotNullRefDefConst         float64 `db:"not_null_ref_def_const"`
}
