package dto

type ConstraintComboNonPk struct {
	NotNullUniqueCheckDefConst string `db:"not_null_unique_check_def_const"`
	NotNullUniqueRefDefConst   string `db:"not_null_unique_ref_def_const"`
	NotNullCheckRefDefConst    string `db:"not_null_check_ref_def_const"`
	NotNullUniqueDefConst      string `db:"not_null_unique_def_const"`
	NotNullCheckDefConst       string `db:"not_null_check_def_const"`
	NotNullRefDefConst         string `db:"not_null_ref_def_const"`
}
