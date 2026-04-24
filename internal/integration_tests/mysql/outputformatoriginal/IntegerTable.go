package dto

import (
	"database/sql"
)

type Integer_table struct {
	I                              sql.NullInt64 `db:"i"`
	Integer_nn                     int           `db:"integer_nn"`
	Integer_nn_unique              int           `db:"integer_nn_unique"`
	Integer_nn_check               int           `db:"integer_nn_check"`
	Integer_unique                 sql.NullInt64 `db:"integer_unique"`
	Integer_unique_check           sql.NullInt64 `db:"integer_unique_check"`
	Integer_unique_ref             sql.NullInt64 `db:"integer_unique_ref"`
	Integer_unique_def_const       sql.NullInt64 `db:"integer_unique_def_const"`
	Integer_unique_def_func        sql.NullInt64 `db:"integer_unique_def_func"`
	Integer_check                  sql.NullInt64 `db:"integer_check"`
	Integer_check_ref              sql.NullInt64 `db:"integer_check_ref"`
	Integer_check_def_const        sql.NullInt64 `db:"integer_check_def_const"`
	Integer_check_def_func         sql.NullInt64 `db:"integer_check_def_func"`
	Integer_ref                    sql.NullInt64 `db:"integer_ref"`
	Integer_ref_unique_check       sql.NullInt64 `db:"integer_ref_unique_check"`
	Integer_def_const              sql.NullInt64 `db:"integer_def_const"`
	Integer_def_const_unique_check sql.NullInt64 `db:"integer_def_const_unique_check"`
	Integer_def_func               sql.NullInt64 `db:"integer_def_func"`
	Integer_def_func_unique_check  sql.NullInt64 `db:"integer_def_func_unique_check"`
}
