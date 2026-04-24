package dto

import (
	"database/sql"
)

type MediumblobTable struct {
	Col                     sql.NullString `db:"col"`
	MediumblobDefConst      sql.NullString `db:"mediumblob_def_const"`
	MediumblobDefFunc       sql.NullString `db:"mediumblob_def_func"`
	MediumblobNn            string         `db:"mediumblob_nn"`
	MediumblobNnCheckCmp    string         `db:"mediumblob_nn_check_cmp"`
	MediumblobNnCheckFn     string         `db:"mediumblob_nn_check_fn"`
	MediumblobNnDefConst    string         `db:"mediumblob_nn_def_const"`
	MediumblobNnDefFunc     string         `db:"mediumblob_nn_def_func"`
	MediumblobCheck         sql.NullString `db:"mediumblob_check"`
	MediumblobCheckDefConst sql.NullString `db:"mediumblob_check_def_const"`
	MediumblobCheckDefFunc  sql.NullString `db:"mediumblob_check_def_func"`
}
