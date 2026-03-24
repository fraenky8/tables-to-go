package dto

type Character struct {
	Character                    *string `db:"character"`
	CharacterCap                 *string `db:"character_cap"`
	CharacterNn                  string  `db:"character_nn"`
	CharacterNnUnique            string  `db:"character_nn_unique"`
	CharacterNnCheckCmp          string  `db:"character_nn_check_cmp"`
	CharacterNnCheckFn           string  `db:"character_nn_check_fn"`
	CharacterNnRef               string  `db:"character_nn_ref"`
	CharacterNnDefConst          string  `db:"character_nn_def_const"`
	CharacterNnDefFunc           string  `db:"character_nn_def_func"`
	CharacterNnUniqueCheck       string  `db:"character_nn_unique_check"`
	CharacterUnique              *string `db:"character_unique"`
	CharacterUniqueCheck         *string `db:"character_unique_check"`
	CharacterUniqueRef           *string `db:"character_unique_ref"`
	CharacterUniqueDefConst      *string `db:"character_unique_def_const"`
	CharacterUniqueDefFunc       *string `db:"character_unique_def_func"`
	CharacterCheck               *string `db:"character_check"`
	CharacterCheckRef            *string `db:"character_check_ref"`
	CharacterCheckDefConst       *string `db:"character_check_def_const"`
	CharacterCheckDefFunc        *string `db:"character_check_def_func"`
	CharacterRef                 *string `db:"character_ref"`
	CharacterRefDefConst         *string `db:"character_ref_def_const"`
	CharacterRefDefFunc          *string `db:"character_ref_def_func"`
	CharacterRefUniqueCheck      *string `db:"character_ref_unique_check"`
	CharacterDefConst            *string `db:"character_def_const"`
	CharacterDefConstUniqueCheck *string `db:"character_def_const_unique_check"`
	CharacterDefFunc             *string `db:"character_def_func"`
	CharacterDefFuncUniqueCheck  *string `db:"character_def_func_unique_check"`
}
