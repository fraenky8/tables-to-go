package dto

import (
	"database/sql"
)

type VarcharTable struct {
	Col                        sql.NullString
	VarcharCap                 sql.NullString
	VarcharNn                  string
	VarcharNnUnique            string
	VarcharNnCheckCmp          string
	VarcharNnCheckFn           string
	VarcharNnRef               string
	VarcharNnDefConst          string
	VarcharNnDefFunc           string
	VarcharNnUniqueCheck       string
	VarcharUnique              sql.NullString
	VarcharUniqueCheck         sql.NullString
	VarcharUniqueRef           sql.NullString
	VarcharUniqueDefConst      sql.NullString
	VarcharUniqueDefFunc       sql.NullString
	VarcharCheck               sql.NullString
	VarcharCheckRef            sql.NullString
	VarcharCheckDefConst       sql.NullString
	VarcharCheckDefFunc        sql.NullString
	VarcharRef                 sql.NullString
	VarcharRefUniqueCheck      sql.NullString
	VarcharDefConst            sql.NullString
	VarcharDefConstUniqueCheck sql.NullString
	VarcharDefFunc             sql.NullString
	VarcharDefFuncUniqueCheck  sql.NullString
}
