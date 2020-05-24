package dto

import (
	"database/sql"
)

type VarbinaryTable struct {
	Col                          sql.NullString `db:"col"`
	VarbinaryCap                 sql.NullString `db:"varbinary_cap"`
	VarbinaryNn                  string         `db:"varbinary_nn"`
	VarbinaryNnUnique            string         `db:"varbinary_nn_unique"`
	VarbinaryNnCheckCmp          string         `db:"varbinary_nn_check_cmp"`
	VarbinaryNnCheckFn           string         `db:"varbinary_nn_check_fn"`
	VarbinaryNnRef               string         `db:"varbinary_nn_ref"`
	VarbinaryNnDefConst          string         `db:"varbinary_nn_def_const"`
	VarbinaryNnDefFunc           string         `db:"varbinary_nn_def_func"`
	VarbinaryNnUniqueCheck       string         `db:"varbinary_nn_unique_check"`
	VarbinaryUnique              sql.NullString `db:"varbinary_unique"`
	VarbinaryUniqueCheck         sql.NullString `db:"varbinary_unique_check"`
	VarbinaryUniqueRef           sql.NullString `db:"varbinary_unique_ref"`
	VarbinaryUniqueDefConst      sql.NullString `db:"varbinary_unique_def_const"`
	VarbinaryUniqueDefFunc       sql.NullString `db:"varbinary_unique_def_func"`
	VarbinaryCheck               sql.NullString `db:"varbinary_check"`
	VarbinaryCheckRef            sql.NullString `db:"varbinary_check_ref"`
	VarbinaryCheckDefConst       sql.NullString `db:"varbinary_check_def_const"`
	VarbinaryCheckDefFunc        sql.NullString `db:"varbinary_check_def_func"`
	VarbinaryRef                 sql.NullString `db:"varbinary_ref"`
	VarbinaryRefUniqueCheck      sql.NullString `db:"varbinary_ref_unique_check"`
	VarbinaryDefConst            sql.NullString `db:"varbinary_def_const"`
	VarbinaryDefConstUniqueCheck sql.NullString `db:"varbinary_def_const_unique_check"`
	VarbinaryDefFunc             sql.NullString `db:"varbinary_def_func"`
	VarbinaryDefFuncUniqueCheck  sql.NullString `db:"varbinary_def_func_unique_check"`
}
