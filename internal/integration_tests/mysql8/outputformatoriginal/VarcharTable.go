package dto

import (
	"database/sql"
)

type Varchar_table struct {
	Col                            sql.NullString `db:"col"`
	Varchar_cap                    sql.NullString `db:"varchar_cap"`
	Varchar_nn                     string         `db:"varchar_nn"`
	Varchar_nn_unique              string         `db:"varchar_nn_unique"`
	Varchar_nn_check_cmp           string         `db:"varchar_nn_check_cmp"`
	Varchar_nn_check_fn            string         `db:"varchar_nn_check_fn"`
	Varchar_nn_ref                 string         `db:"varchar_nn_ref"`
	Varchar_nn_def_const           string         `db:"varchar_nn_def_const"`
	Varchar_nn_def_func            string         `db:"varchar_nn_def_func"`
	Varchar_nn_unique_check        string         `db:"varchar_nn_unique_check"`
	Varchar_unique                 sql.NullString `db:"varchar_unique"`
	Varchar_unique_check           sql.NullString `db:"varchar_unique_check"`
	Varchar_unique_ref             sql.NullString `db:"varchar_unique_ref"`
	Varchar_unique_def_const       sql.NullString `db:"varchar_unique_def_const"`
	Varchar_unique_def_func        sql.NullString `db:"varchar_unique_def_func"`
	Varchar_check                  sql.NullString `db:"varchar_check"`
	Varchar_check_ref              sql.NullString `db:"varchar_check_ref"`
	Varchar_check_def_const        sql.NullString `db:"varchar_check_def_const"`
	Varchar_check_def_func         sql.NullString `db:"varchar_check_def_func"`
	Varchar_ref                    sql.NullString `db:"varchar_ref"`
	Varchar_ref_unique_check       sql.NullString `db:"varchar_ref_unique_check"`
	Varchar_def_const              sql.NullString `db:"varchar_def_const"`
	Varchar_def_const_unique_check sql.NullString `db:"varchar_def_const_unique_check"`
	Varchar_def_func               sql.NullString `db:"varchar_def_func"`
	Varchar_def_func_unique_check  sql.NullString `db:"varchar_def_func_unique_check"`
}
