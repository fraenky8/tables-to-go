package dto

import (
	"database/sql"
)

type IntegerTable struct {
	I                          sql.NullInt64
	IntegerNn                  int
	IntegerNnUnique            int
	IntegerNnCheck             int
	IntegerUnique              sql.NullInt64
	IntegerUniqueCheck         sql.NullInt64
	IntegerUniqueRef           sql.NullInt64
	IntegerUniqueDefConst      sql.NullInt64
	IntegerUniqueDefFunc       sql.NullInt64
	IntegerCheck               sql.NullInt64
	IntegerCheckRef            sql.NullInt64
	IntegerCheckDefConst       sql.NullInt64
	IntegerCheckDefFunc        sql.NullInt64
	IntegerRef                 sql.NullInt64
	IntegerRefUniqueCheck      sql.NullInt64
	IntegerDefConst            sql.NullInt64
	IntegerDefConstUniqueCheck sql.NullInt64
	IntegerDefFunc             sql.NullInt64
	IntegerDefFuncUniqueCheck  sql.NullInt64
}
