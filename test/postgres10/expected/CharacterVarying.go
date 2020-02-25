package dto

import (
	"database/sql"
)

type CharacterVarying struct {
	CharacterVarying                    sql.NullString `db:"character_varying"`
	CharacterVaryingCap                 sql.NullString `db:"character_varying_cap"`
	CharacterVaryingNn                  string         `db:"character_varying_nn"`
	CharacterVaryingNnUnique            string         `db:"character_varying_nn_unique"`
	CharacterVaryingNnCheckCmp          string         `db:"character_varying_nn_check_cmp"`
	CharacterVaryingNnCheckFn           string         `db:"character_varying_nn_check_fn"`
	CharacterVaryingNnRef               string         `db:"character_varying_nn_ref"`
	CharacterVaryingNnDefConst          string         `db:"character_varying_nn_def_const"`
	CharacterVaryingNnDefFunc           string         `db:"character_varying_nn_def_func"`
	CharacterVaryingNnUniqueCheck       string         `db:"character_varying_nn_unique_check"`
	CharacterVaryingUnique              sql.NullString `db:"character_varying_unique"`
	CharacterVaryingUniqueCheck         sql.NullString `db:"character_varying_unique_check"`
	CharacterVaryingUniqueRef           sql.NullString `db:"character_varying_unique_ref"`
	CharacterVaryingUniqueDefConst      sql.NullString `db:"character_varying_unique_def_const"`
	CharacterVaryingUniqueDefFunc       sql.NullString `db:"character_varying_unique_def_func"`
	CharacterVaryingCheck               sql.NullString `db:"character_varying_check"`
	CharacterVaryingCheckRef            sql.NullString `db:"character_varying_check_ref"`
	CharacterVaryingCheckDefConst       sql.NullString `db:"character_varying_check_def_const"`
	CharacterVaryingCheckDefFunc        sql.NullString `db:"character_varying_check_def_func"`
	CharacterVaryingRef                 sql.NullString `db:"character_varying_ref"`
	CharacterVaryingRefDefConst         sql.NullString `db:"character_varying_ref_def_const"`
	CharacterVaryingRefDefFunc          sql.NullString `db:"character_varying_ref_def_func"`
	CharacterVaryingRefUniqueCheck      sql.NullString `db:"character_varying_ref_unique_check"`
	CharacterVaryingDefConst            sql.NullString `db:"character_varying_def_const"`
	CharacterVaryingDefConstUniqueCheck sql.NullString `db:"character_varying_def_const_unique_check"`
	CharacterVaryingDefFunc             sql.NullString `db:"character_varying_def_func"`
	CharacterVaryingDefFuncUniqueCheck  sql.NullString `db:"character_varying_def_func_unique_check"`
}
