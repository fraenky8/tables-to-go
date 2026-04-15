package dto

type TextTable struct {
	T            *string `db:"t"`
	TextNn       string  `db:"text_nn"`
	TextUnique   *string `db:"text_unique"`
	TextCheck    *string `db:"text_check"`
	TextRef      *string `db:"text_ref"`
	TextDefConst *string `db:"text_def_const"`
	TextPk       *string `db:"text_pk"`
}
