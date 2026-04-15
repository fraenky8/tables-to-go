package dto

type StrictTypes struct {
	StrictID       *string `db:"strict_id"`
	StrictInt      *string `db:"strict_int"`
	StrictReal     *string `db:"strict_real"`
	StrictText     *string `db:"strict_text"`
	StrictBlob     *string `db:"strict_blob"`
	StrictAny      *string `db:"strict_any"`
	StrictNn       string  `db:"strict_nn"`
	StrictUnique   *string `db:"strict_unique"`
	StrictCheck    *string `db:"strict_check"`
	StrictRef      *string `db:"strict_ref"`
	StrictDefConst *string `db:"strict_def_const"`
}
