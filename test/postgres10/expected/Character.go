package dto

import (
	"database/sql"
)

type Character struct {
	Character                    sql.NullString `db:"character"`
	CharacterCap                 sql.NullString `db:"character_cap"`
	CharacterNn                  string         `db:"character_nn"`
	CharacterNnUnique            string         `db:"character_nn_unique"`
	CharacterNnCheckCmp          string         `db:"character_nn_check_cmp"`
	CharacterNnCheckFn           string         `db:"character_nn_check_fn"`
	CharacterNnRef               string         `db:"character_nn_ref"`
	CharacterNnDefConst          string         `db:"character_nn_def_const"`
	CharacterNnDefFunc           string         `db:"character_nn_def_func"`
	CharacterNnUniqueCheck       string         `db:"character_nn_unique_check"`
	CharacterUnique              sql.NullString `db:"character_unique"`
	CharacterUniqueCheck         sql.NullString `db:"character_unique_check"`
	CharacterUniqueRef           sql.NullString `db:"character_unique_ref"`
	CharacterUniqueDefConst      sql.NullString `db:"character_unique_def_const"`
	CharacterUniqueDefFunc       sql.NullString `db:"character_unique_def_func"`
	CharacterCheck               sql.NullString `db:"character_check"`
	CharacterCheckRef            sql.NullString `db:"character_check_ref"`
	CharacterCheckDefConst       sql.NullString `db:"character_check_def_const"`
	CharacterCheckDefFunc        sql.NullString `db:"character_check_def_func"`
	CharacterRef                 sql.NullString `db:"character_ref"`
	CharacterRefDefConst         sql.NullString `db:"character_ref_def_const"`
	CharacterRefDefFunc          sql.NullString `db:"character_ref_def_func"`
	CharacterRefUniqueCheck      sql.NullString `db:"character_ref_unique_check"`
	CharacterDefConst            sql.NullString `db:"character_def_const"`
	CharacterDefConstUniqueCheck sql.NullString `db:"character_def_const_unique_check"`
	CharacterDefFunc             sql.NullString `db:"character_def_func"`
	CharacterDefFuncUniqueCheck  sql.NullString `db:"character_def_func_unique_check"`
}
