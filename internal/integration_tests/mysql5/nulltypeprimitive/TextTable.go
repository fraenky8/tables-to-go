package dto

type TextTable struct {
	Col               *string `db:"col"`
	TextDefConst      *string `db:"text_def_const"`
	TextDefFunc       *string `db:"text_def_func"`
	TextRef           *string `db:"text_ref"`
	TextNn            string  `db:"text_nn"`
	TextNnCheckCmp    string  `db:"text_nn_check_cmp"`
	TextNnCheckFn     string  `db:"text_nn_check_fn"`
	TextNnRef         string  `db:"text_nn_ref"`
	TextNnDefConst    string  `db:"text_nn_def_const"`
	TextNnDefFunc     string  `db:"text_nn_def_func"`
	TextCheck         *string `db:"text_check"`
	TextCheckRef      *string `db:"text_check_ref"`
	TextCheckDefConst *string `db:"text_check_def_const"`
	TextCheckDefFunc  *string `db:"text_check_def_func"`
}
