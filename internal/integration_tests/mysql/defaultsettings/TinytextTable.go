package dto

import (
	"database/sql"
)

type TinytextTable struct {
	Col                   sql.NullString `db:"col"`
	TinytextDefConst      sql.NullString `db:"tinytext_def_const"`
	TinytextDefFunc       sql.NullString `db:"tinytext_def_func"`
	TinytextNn            string         `db:"tinytext_nn"`
	TinytextNnCheckCmp    string         `db:"tinytext_nn_check_cmp"`
	TinytextNnCheckFn     string         `db:"tinytext_nn_check_fn"`
	TinytextNnDefConst    string         `db:"tinytext_nn_def_const"`
	TinytextNnDefFunc     string         `db:"tinytext_nn_def_func"`
	TinytextCheck         sql.NullString `db:"tinytext_check"`
	TinytextCheckDefConst sql.NullString `db:"tinytext_check_def_const"`
	TinytextCheckDefFunc  sql.NullString `db:"tinytext_check_def_func"`
}
