package dto

import (
	"database/sql"
)

type MediumblobTable struct {
	Col                     sql.NullString `db:"col"`
	MediumblobDefConst      sql.NullString `db:"mediumblob_def_const"`
	MediumblobDefFunc       sql.NullString `db:"mediumblob_def_func"`
	MediumblobRef           sql.NullString `db:"mediumblob_ref"`
	MediumblobNn            string         `db:"mediumblob_nn"`
	MediumblobNnCheckCmp    string         `db:"mediumblob_nn_check_cmp"`
	MediumblobNnCheckFn     string         `db:"mediumblob_nn_check_fn"`
	MediumblobNnRef         string         `db:"mediumblob_nn_ref"`
	MediumblobNnDefConst    string         `db:"mediumblob_nn_def_const"`
	MediumblobNnDefFunc     string         `db:"mediumblob_nn_def_func"`
	MediumblobCheck         sql.NullString `db:"mediumblob_check"`
	MediumblobCheckRef      sql.NullString `db:"mediumblob_check_ref"`
	MediumblobCheckDefConst sql.NullString `db:"mediumblob_check_def_const"`
	MediumblobCheckDefFunc  sql.NullString `db:"mediumblob_check_def_func"`
}
