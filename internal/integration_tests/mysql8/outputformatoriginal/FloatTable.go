package dto

import (
	"database/sql"
)

type Float_table struct {
	Col                          sql.NullFloat64 `db:"col"`
	Float_nn                     float64         `db:"float_nn"`
	Float_nn_unique              float64         `db:"float_nn_unique"`
	Float_nn_check               float64         `db:"float_nn_check"`
	Float_nn_ref                 float64         `db:"float_nn_ref"`
	Float_nn_def_const           float64         `db:"float_nn_def_const"`
	Float_nn_def_func            float64         `db:"float_nn_def_func"`
	Float_nn_unique_check        float64         `db:"float_nn_unique_check"`
	Float_unique                 sql.NullFloat64 `db:"float_unique"`
	Float_unique_check           sql.NullFloat64 `db:"float_unique_check"`
	Float_unique_ref             sql.NullFloat64 `db:"float_unique_ref"`
	Float_unique_def_const       sql.NullFloat64 `db:"float_unique_def_const"`
	Float_unique_def_func        sql.NullFloat64 `db:"float_unique_def_func"`
	Float_check                  sql.NullFloat64 `db:"float_check"`
	Float_check_ref              sql.NullFloat64 `db:"float_check_ref"`
	Float_check_def_const        sql.NullFloat64 `db:"float_check_def_const"`
	Float_check_def_func         sql.NullFloat64 `db:"float_check_def_func"`
	Float_ref                    sql.NullFloat64 `db:"float_ref"`
	Float_ref_unique_check       sql.NullFloat64 `db:"float_ref_unique_check"`
	Float_def_const              sql.NullFloat64 `db:"float_def_const"`
	Float_def_const_unique_check sql.NullFloat64 `db:"float_def_const_unique_check"`
	Float_def_func               sql.NullFloat64 `db:"float_def_func"`
	Float_def_func_unique_check  sql.NullFloat64 `db:"float_def_func_unique_check"`
}
