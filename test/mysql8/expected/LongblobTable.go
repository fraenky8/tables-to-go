package dto

import (
	"database/sql"
)

type LongblobTable struct {
	Col                   sql.NullString `db:"col"`
	LongblobDefConst      sql.NullString `db:"longblob_def_const"`
	LongblobDefFunc       sql.NullString `db:"longblob_def_func"`
	LongblobRef           sql.NullString `db:"longblob_ref"`
	LongblobNn            string         `db:"longblob_nn"`
	LongblobNnCheckCmp    string         `db:"longblob_nn_check_cmp"`
	LongblobNnCheckFn     string         `db:"longblob_nn_check_fn"`
	LongblobNnRef         string         `db:"longblob_nn_ref"`
	LongblobNnDefConst    string         `db:"longblob_nn_def_const"`
	LongblobNnDefFunc     string         `db:"longblob_nn_def_func"`
	LongblobCheck         sql.NullString `db:"longblob_check"`
	LongblobCheckRef      sql.NullString `db:"longblob_check_ref"`
	LongblobCheckDefConst sql.NullString `db:"longblob_check_def_const"`
	LongblobCheckDefFunc  sql.NullString `db:"longblob_check_def_func"`
}
