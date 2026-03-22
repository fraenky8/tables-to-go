package dto

import (
	"database/sql"
)

type TextTable struct {
	Col               sql.NullString `db:"col"`
	TextDefConst      sql.NullString `db:"text_def_const"`
	TextDefFunc       sql.NullString `db:"text_def_func"`
	TextRef           sql.NullString `db:"text_ref"`
	TextNn            string         `db:"text_nn"`
	TextNnCheckCmp    string         `db:"text_nn_check_cmp"`
	TextNnCheckFn     string         `db:"text_nn_check_fn"`
	TextNnRef         string         `db:"text_nn_ref"`
	TextNnDefConst    string         `db:"text_nn_def_const"`
	TextNnDefFunc     string         `db:"text_nn_def_func"`
	TextCheck         sql.NullString `db:"text_check"`
	TextCheckRef      sql.NullString `db:"text_check_ref"`
	TextCheckDefConst sql.NullString `db:"text_check_def_const"`
	TextCheckDefFunc  sql.NullString `db:"text_check_def_func"`
}
