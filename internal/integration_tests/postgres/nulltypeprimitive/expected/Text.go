package dto

type Text struct {
	Text                    *string `db:"text"`
	TextNn                  string  `db:"text_nn"`
	TextNnUnique            string  `db:"text_nn_unique"`
	TextNnCheckCmp          string  `db:"text_nn_check_cmp"`
	TextNnCheckFn           string  `db:"text_nn_check_fn"`
	TextNnRef               string  `db:"text_nn_ref"`
	TextNnDefConst          string  `db:"text_nn_def_const"`
	TextNnDefFunc           string  `db:"text_nn_def_func"`
	TextNnUniqueCheck       string  `db:"text_nn_unique_check"`
	TextUnique              *string `db:"text_unique"`
	TextUniqueCheck         *string `db:"text_unique_check"`
	TextUniqueRef           *string `db:"text_unique_ref"`
	TextUniqueDefConst      *string `db:"text_unique_def_const"`
	TextUniqueDefFunc       *string `db:"text_unique_def_func"`
	TextCheck               *string `db:"text_check"`
	TextCheckRef            *string `db:"text_check_ref"`
	TextCheckDefConst       *string `db:"text_check_def_const"`
	TextCheckDefFunc        *string `db:"text_check_def_func"`
	TextRef                 *string `db:"text_ref"`
	TextRefDefConst         *string `db:"text_ref_def_const"`
	TextRefDefFunc          *string `db:"text_ref_def_func"`
	TextRefUniqueCheck      *string `db:"text_ref_unique_check"`
	TextDefConst            *string `db:"text_def_const"`
	TextDefConstUniqueCheck *string `db:"text_def_const_unique_check"`
	TextDefFunc             *string `db:"text_def_func"`
	TextDefFuncUniqueCheck  *string `db:"text_def_func_unique_check"`
}
