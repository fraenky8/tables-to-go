package dto

type TextTable struct {
	Col               *string `db:"col"`
	TextDefConst      *string `db:"text_def_const"`
	TextDefFunc       *string `db:"text_def_func"`
	TextNn            string  `db:"text_nn"`
	TextNnCheckCmp    string  `db:"text_nn_check_cmp"`
	TextNnCheckFn     string  `db:"text_nn_check_fn"`
	TextNnDefConst    string  `db:"text_nn_def_const"`
	TextNnDefFunc     string  `db:"text_nn_def_func"`
	TextCheck         *string `db:"text_check"`
	TextCheckDefConst *string `db:"text_check_def_const"`
	TextCheckDefFunc  *string `db:"text_check_def_func"`
}
