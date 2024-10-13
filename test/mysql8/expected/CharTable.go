package dto

import (
	"database/sql"
)

type CharTable struct {
	Col                     sql.NullString `db:"col"`
	CharCap                 sql.NullString `db:"char_cap"`
	CharNn                  string         `db:"char_nn"`
	CharNnUnique            string         `db:"char_nn_unique"`
	CharNnCheckCmp          string         `db:"char_nn_check_cmp"`
	CharNnCheckFn           string         `db:"char_nn_check_fn"`
	CharNnRef               string         `db:"char_nn_ref"`
	CharNnDefConst          string         `db:"char_nn_def_const"`
	CharNnDefFunc           string         `db:"char_nn_def_func"`
	CharNnUniqueCheck       string         `db:"char_nn_unique_check"`
	CharUnique              sql.NullString `db:"char_unique"`
	CharUniqueCheck         sql.NullString `db:"char_unique_check"`
	CharUniqueRef           sql.NullString `db:"char_unique_ref"`
	CharUniqueDefConst      sql.NullString `db:"char_unique_def_const"`
	CharUniqueDefFunc       sql.NullString `db:"char_unique_def_func"`
	CharCheck               sql.NullString `db:"char_check"`
	CharCheckRef            sql.NullString `db:"char_check_ref"`
	CharCheckDefConst       sql.NullString `db:"char_check_def_const"`
	CharCheckDefFunc        sql.NullString `db:"char_check_def_func"`
	CharRef                 sql.NullString `db:"char_ref"`
	CharRefUniqueCheck      sql.NullString `db:"char_ref_unique_check"`
	CharDefConst            sql.NullString `db:"char_def_const"`
	CharDefConstUniqueCheck sql.NullString `db:"char_def_const_unique_check"`
	CharDefFunc             sql.NullString `db:"char_def_func"`
	CharDefFuncUniqueCheck  sql.NullString `db:"char_def_func_unique_check"`
}
