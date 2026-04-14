package dto

import (
	"database/sql"
	"time"
)

type Datetime_table struct {
	Datetime                        sql.NullTime `db:"datetime"`
	Datetime_nn                     time.Time    `db:"datetime_nn"`
	Datetime_nn_unique              time.Time    `db:"datetime_nn_unique"`
	Datetime_nn_check               time.Time    `db:"datetime_nn_check"`
	Datetime_nn_ref                 time.Time    `db:"datetime_nn_ref"`
	Datetime_nn_def_const           time.Time    `db:"datetime_nn_def_const"`
	Datetime_nn_def_func            time.Time    `db:"datetime_nn_def_func"`
	Datetime_nn_unique_check        time.Time    `db:"datetime_nn_unique_check"`
	Datetime_unique                 sql.NullTime `db:"datetime_unique"`
	Datetime_unique_check           sql.NullTime `db:"datetime_unique_check"`
	Datetime_unique_ref             sql.NullTime `db:"datetime_unique_ref"`
	Datetime_unique_def_const       sql.NullTime `db:"datetime_unique_def_const"`
	Datetime_unique_def_func        sql.NullTime `db:"datetime_unique_def_func"`
	Datetime_check                  sql.NullTime `db:"datetime_check"`
	Datetime_check_ref              sql.NullTime `db:"datetime_check_ref"`
	Datetime_check_def_const        sql.NullTime `db:"datetime_check_def_const"`
	Datetime_check_def_func         sql.NullTime `db:"datetime_check_def_func"`
	Datetime_ref                    sql.NullTime `db:"datetime_ref"`
	Datetime_ref_unique_check       sql.NullTime `db:"datetime_ref_unique_check"`
	Datetime_def_const              sql.NullTime `db:"datetime_def_const"`
	Datetime_def_const_unique_check sql.NullTime `db:"datetime_def_const_unique_check"`
	Datetime_def_func               sql.NullTime `db:"datetime_def_func"`
	Datetime_def_func_unique_check  sql.NullTime `db:"datetime_def_func_unique_check"`
}
