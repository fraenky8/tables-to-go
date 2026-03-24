package models

type ConstraintComboNonPk struct {
	NotNullUniqueCheckDefConst float64 `db:"not_null_unique_check_def_const"`
	NotNullUniqueCheckDefFunc  float64 `db:"not_null_unique_check_def_func"`
	NotNullUniqueRefDefConst   float64 `db:"not_null_unique_ref_def_const"`
	NotNullUniqueRefDefFunc    float64 `db:"not_null_unique_ref_def_func"`
	NotNullCheckRefDefConst    float64 `db:"not_null_check_ref_def_const"`
	NotNullCheckRefDefFunc     float64 `db:"not_null_check_ref_def_func"`
	NotNullUniqueDefConst      float64 `db:"not_null_unique_def_const"`
	NotNullUniqueDefFunc       float64 `db:"not_null_unique_def_func"`
	NotNullCheckDefConst       float64 `db:"not_null_check_def_const"`
	NotNullCheckDefFunc        float64 `db:"not_null_check_def_func"`
	NotNullRefDefConst         float64 `db:"not_null_ref_def_const"`
	NotNullRefDefFunc          float64 `db:"not_null_ref_def_func"`
}
