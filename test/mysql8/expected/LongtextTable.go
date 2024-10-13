package dto

import (
	"database/sql"
)

type LongtextTable struct {
	Col                   sql.NullString `db:"col"`
	LongtextDefConst      sql.NullString `db:"longtext_def_const"`
	LongtextDefFunc       sql.NullString `db:"longtext_def_func"`
	LongtextRef           sql.NullString `db:"longtext_ref"`
	LongtextNn            string         `db:"longtext_nn"`
	LongtextNnCheckCmp    string         `db:"longtext_nn_check_cmp"`
	LongtextNnCheckFn     string         `db:"longtext_nn_check_fn"`
	LongtextNnRef         string         `db:"longtext_nn_ref"`
	LongtextNnDefConst    string         `db:"longtext_nn_def_const"`
	LongtextNnDefFunc     string         `db:"longtext_nn_def_func"`
	LongtextCheck         sql.NullString `db:"longtext_check"`
	LongtextCheckRef      sql.NullString `db:"longtext_check_ref"`
	LongtextCheckDefConst sql.NullString `db:"longtext_check_def_const"`
	LongtextCheckDefFunc  sql.NullString `db:"longtext_check_def_func"`
}
