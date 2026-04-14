package dto

type CharacterVarying struct {
	CharacterVarying                    *string `db:"character_varying"`
	CharacterVaryingCap                 *string `db:"character_varying_cap"`
	CharacterVaryingNn                  string  `db:"character_varying_nn"`
	CharacterVaryingNnUnique            string  `db:"character_varying_nn_unique"`
	CharacterVaryingNnCheckCmp          string  `db:"character_varying_nn_check_cmp"`
	CharacterVaryingNnCheckFn           string  `db:"character_varying_nn_check_fn"`
	CharacterVaryingNnRef               string  `db:"character_varying_nn_ref"`
	CharacterVaryingNnDefConst          string  `db:"character_varying_nn_def_const"`
	CharacterVaryingNnDefFunc           string  `db:"character_varying_nn_def_func"`
	CharacterVaryingNnUniqueCheck       string  `db:"character_varying_nn_unique_check"`
	CharacterVaryingUnique              *string `db:"character_varying_unique"`
	CharacterVaryingUniqueCheck         *string `db:"character_varying_unique_check"`
	CharacterVaryingUniqueRef           *string `db:"character_varying_unique_ref"`
	CharacterVaryingUniqueDefConst      *string `db:"character_varying_unique_def_const"`
	CharacterVaryingUniqueDefFunc       *string `db:"character_varying_unique_def_func"`
	CharacterVaryingCheck               *string `db:"character_varying_check"`
	CharacterVaryingCheckRef            *string `db:"character_varying_check_ref"`
	CharacterVaryingCheckDefConst       *string `db:"character_varying_check_def_const"`
	CharacterVaryingCheckDefFunc        *string `db:"character_varying_check_def_func"`
	CharacterVaryingRef                 *string `db:"character_varying_ref"`
	CharacterVaryingRefDefConst         *string `db:"character_varying_ref_def_const"`
	CharacterVaryingRefDefFunc          *string `db:"character_varying_ref_def_func"`
	CharacterVaryingRefUniqueCheck      *string `db:"character_varying_ref_unique_check"`
	CharacterVaryingDefConst            *string `db:"character_varying_def_const"`
	CharacterVaryingDefConstUniqueCheck *string `db:"character_varying_def_const_unique_check"`
	CharacterVaryingDefFunc             *string `db:"character_varying_def_func"`
	CharacterVaryingDefFuncUniqueCheck  *string `db:"character_varying_def_func_unique_check"`
}
