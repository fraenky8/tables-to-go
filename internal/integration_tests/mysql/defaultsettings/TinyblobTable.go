package dto

import (
	"database/sql"
)

type TinyblobTable struct {
	Col                   sql.NullString `db:"col"`
	TinyblobDefConst      sql.NullString `db:"tinyblob_def_const"`
	TinyblobDefFunc       sql.NullString `db:"tinyblob_def_func"`
	TinyblobNn            string         `db:"tinyblob_nn"`
	TinyblobNnCheckCmp    string         `db:"tinyblob_nn_check_cmp"`
	TinyblobNnCheckFn     string         `db:"tinyblob_nn_check_fn"`
	TinyblobNnDefConst    string         `db:"tinyblob_nn_def_const"`
	TinyblobNnDefFunc     string         `db:"tinyblob_nn_def_func"`
	TinyblobCheck         sql.NullString `db:"tinyblob_check"`
	TinyblobCheckDefConst sql.NullString `db:"tinyblob_check_def_const"`
	TinyblobCheckDefFunc  sql.NullString `db:"tinyblob_check_def_func"`
}
