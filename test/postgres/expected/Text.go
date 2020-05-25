package dto

import (
	"database/sql"
)

type Text struct {
	Text                    sql.NullString `db:"text"`
	TextNn                  string         `db:"text_nn"`
	TextNnUnique            string         `db:"text_nn_unique"`
	TextNnCheckCmp          string         `db:"text_nn_check_cmp"`
	TextNnCheckFn           string         `db:"text_nn_check_fn"`
	TextNnRef               string         `db:"text_nn_ref"`
	TextNnDefConst          string         `db:"text_nn_def_const"`
	TextNnDefFunc           string         `db:"text_nn_def_func"`
	TextNnUniqueCheck       string         `db:"text_nn_unique_check"`
	TextUnique              sql.NullString `db:"text_unique"`
	TextUniqueCheck         sql.NullString `db:"text_unique_check"`
	TextUniqueRef           sql.NullString `db:"text_unique_ref"`
	TextUniqueDefConst      sql.NullString `db:"text_unique_def_const"`
	TextUniqueDefFunc       sql.NullString `db:"text_unique_def_func"`
	TextCheck               sql.NullString `db:"text_check"`
	TextCheckRef            sql.NullString `db:"text_check_ref"`
	TextCheckDefConst       sql.NullString `db:"text_check_def_const"`
	TextCheckDefFunc        sql.NullString `db:"text_check_def_func"`
	TextRef                 sql.NullString `db:"text_ref"`
	TextRefDefConst         sql.NullString `db:"text_ref_def_const"`
	TextRefDefFunc          sql.NullString `db:"text_ref_def_func"`
	TextRefUniqueCheck      sql.NullString `db:"text_ref_unique_check"`
	TextDefConst            sql.NullString `db:"text_def_const"`
	TextDefConstUniqueCheck sql.NullString `db:"text_def_const_unique_check"`
	TextDefFunc             sql.NullString `db:"text_def_func"`
	TextDefFuncUniqueCheck  sql.NullString `db:"text_def_func_unique_check"`
}
