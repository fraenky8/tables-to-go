package dto

import (
	"database/sql"
)

type LongtextTable struct {
	Col                   sql.NullString `db:"col"`
	LongtextDefConst      sql.NullString `db:"longtext_def_const"`
	LongtextDefFunc       sql.NullString `db:"longtext_def_func"`
	LongtextNn            string         `db:"longtext_nn"`
	LongtextNnCheckCmp    string         `db:"longtext_nn_check_cmp"`
	LongtextNnCheckFn     string         `db:"longtext_nn_check_fn"`
	LongtextNnDefConst    string         `db:"longtext_nn_def_const"`
	LongtextNnDefFunc     string         `db:"longtext_nn_def_func"`
	LongtextCheck         sql.NullString `db:"longtext_check"`
	LongtextCheckDefConst sql.NullString `db:"longtext_check_def_const"`
	LongtextCheckDefFunc  sql.NullString `db:"longtext_check_def_func"`
}
