package dto

import (
	"database/sql"
)

type TinytextTable struct {
	Col                   sql.NullString `db:"col"`
	TinytextDefConst      sql.NullString `db:"tinytext_def_const"`
	TinytextDefFunc       sql.NullString `db:"tinytext_def_func"`
	TinytextRef           sql.NullString `db:"tinytext_ref"`
	TinytextNn            string         `db:"tinytext_nn"`
	TinytextNnCheckCmp    string         `db:"tinytext_nn_check_cmp"`
	TinytextNnCheckFn     string         `db:"tinytext_nn_check_fn"`
	TinytextNnRef         string         `db:"tinytext_nn_ref"`
	TinytextNnDefConst    string         `db:"tinytext_nn_def_const"`
	TinytextNnDefFunc     string         `db:"tinytext_nn_def_func"`
	TinytextCheck         sql.NullString `db:"tinytext_check"`
	TinytextCheckRef      sql.NullString `db:"tinytext_check_ref"`
	TinytextCheckDefConst sql.NullString `db:"tinytext_check_def_const"`
	TinytextCheckDefFunc  sql.NullString `db:"tinytext_check_def_func"`
}
