package dto

type DoublePrecisionTable struct {
	Col                                *float64 `db:"col"`
	DoublePrecisionNn                  float64  `db:"double_precision_nn"`
	DoublePrecisionNnUnique            float64  `db:"double_precision_nn_unique"`
	DoublePrecisionNnCheck             float64  `db:"double_precision_nn_check"`
	DoublePrecisionNnRef               float64  `db:"double_precision_nn_ref"`
	DoublePrecisionNnDefConst          float64  `db:"double_precision_nn_def_const"`
	DoublePrecisionNnDefFunc           float64  `db:"double_precision_nn_def_func"`
	DoublePrecisionNnUniqueCheck       float64  `db:"double_precision_nn_unique_check"`
	DoublePrecisionUnique              *float64 `db:"double_precision_unique"`
	DoublePrecisionUniqueCheck         *float64 `db:"double_precision_unique_check"`
	DoublePrecisionUniqueRef           *float64 `db:"double_precision_unique_ref"`
	DoublePrecisionUniqueDefConst      *float64 `db:"double_precision_unique_def_const"`
	DoublePrecisionUniqueDefFunc       *float64 `db:"double_precision_unique_def_func"`
	DoublePrecisionCheck               *float64 `db:"double_precision_check"`
	DoublePrecisionCheckRef            *float64 `db:"double_precision_check_ref"`
	DoublePrecisionCheckDefConst       *float64 `db:"double_precision_check_def_const"`
	DoublePrecisionCheckDefFunc        *float64 `db:"double_precision_check_def_func"`
	DoublePrecisionRef                 *float64 `db:"double_precision_ref"`
	DoublePrecisionRefUniqueCheck      *float64 `db:"double_precision_ref_unique_check"`
	DoublePrecisionDefConst            *float64 `db:"double_precision_def_const"`
	DoublePrecisionDefConstUniqueCheck *float64 `db:"double_precision_def_const_unique_check"`
	DoublePrecisionDefFunc             *float64 `db:"double_precision_def_func"`
	DoublePrecisionDefFuncUniqueCheck  *float64 `db:"double_precision_def_func_unique_check"`
}
