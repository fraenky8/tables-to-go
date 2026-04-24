package dto

import (
	"database/sql"
	"time"
)

type DatetimeTable struct {
	Datetime                    sql.NullTime
	DatetimeNn                  time.Time
	DatetimeNnUnique            time.Time
	DatetimeNnCheck             time.Time
	DatetimeNnRef               time.Time
	DatetimeNnDefConst          time.Time
	DatetimeNnDefFunc           time.Time
	DatetimeNnUniqueCheck       time.Time
	DatetimeUnique              sql.NullTime
	DatetimeUniqueCheck         sql.NullTime
	DatetimeUniqueRef           sql.NullTime
	DatetimeUniqueDefConst      sql.NullTime
	DatetimeUniqueDefFunc       sql.NullTime
	DatetimeCheck               sql.NullTime
	DatetimeCheckRef            sql.NullTime
	DatetimeCheckDefConst       sql.NullTime
	DatetimeCheckDefFunc        sql.NullTime
	DatetimeRef                 sql.NullTime
	DatetimeRefUniqueCheck      sql.NullTime
	DatetimeDefConst            sql.NullTime
	DatetimeDefConstUniqueCheck sql.NullTime
	DatetimeDefFunc             sql.NullTime
	DatetimeDefFuncUniqueCheck  sql.NullTime
}
