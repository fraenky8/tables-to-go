package dto

type StrictTypes struct {
	StrictID       *int     `db:"strict_id"`
	StrictInt      *int     `db:"strict_int"`
	StrictReal     *float64 `db:"strict_real"`
	StrictText     *string  `db:"strict_text"`
	StrictBlob     *string  `db:"strict_blob"`
	StrictAny      *string  `db:"strict_any"`
	StrictNn       string   `db:"strict_nn"`
	StrictUnique   *float64 `db:"strict_unique"`
	StrictCheck    *int     `db:"strict_check"`
	StrictRef      *int     `db:"strict_ref"`
	StrictDefConst *string  `db:"strict_def_const"`
}
