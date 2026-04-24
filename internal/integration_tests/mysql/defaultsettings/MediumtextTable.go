package dto

import (
	"database/sql"
)

type MediumtextTable struct {
	Col                     sql.NullString `db:"col"`
	MediumtextDefConst      sql.NullString `db:"mediumtext_def_const"`
	MediumtextDefFunc       sql.NullString `db:"mediumtext_def_func"`
	MediumtextNn            string         `db:"mediumtext_nn"`
	MediumtextNnCheckCmp    string         `db:"mediumtext_nn_check_cmp"`
	MediumtextNnCheckFn     string         `db:"mediumtext_nn_check_fn"`
	MediumtextNnDefConst    string         `db:"mediumtext_nn_def_const"`
	MediumtextNnDefFunc     string         `db:"mediumtext_nn_def_func"`
	MediumtextCheck         sql.NullString `db:"mediumtext_check"`
	MediumtextCheckDefConst sql.NullString `db:"mediumtext_check_def_const"`
	MediumtextCheckDefFunc  sql.NullString `db:"mediumtext_check_def_func"`
}
