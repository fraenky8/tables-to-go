package dto

import (
	"database/sql"
)

type TinyblobTable struct {
	Col                   sql.NullString `db:"col"`
	TinyblobDefConst      sql.NullString `db:"tinyblob_def_const"`
	TinyblobDefFunc       sql.NullString `db:"tinyblob_def_func"`
	TinyblobRef           sql.NullString `db:"tinyblob_ref"`
	TinyblobNn            string         `db:"tinyblob_nn"`
	TinyblobNnCheckCmp    string         `db:"tinyblob_nn_check_cmp"`
	TinyblobNnCheckFn     string         `db:"tinyblob_nn_check_fn"`
	TinyblobNnRef         string         `db:"tinyblob_nn_ref"`
	TinyblobNnDefConst    string         `db:"tinyblob_nn_def_const"`
	TinyblobNnDefFunc     string         `db:"tinyblob_nn_def_func"`
	TinyblobCheck         sql.NullString `db:"tinyblob_check"`
	TinyblobCheckRef      sql.NullString `db:"tinyblob_check_ref"`
	TinyblobCheckDefConst sql.NullString `db:"tinyblob_check_def_const"`
	TinyblobCheckDefFunc  sql.NullString `db:"tinyblob_check_def_func"`
}
