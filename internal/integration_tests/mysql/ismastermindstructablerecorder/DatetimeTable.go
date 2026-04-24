package dto

import (
	"database/sql"
	"time"

	"github.com/Masterminds/structable"
)

type DatetimeTable struct {
	structable.Recorder

	Datetime                    sql.NullTime `stbl:"datetime"`
	DatetimeNn                  time.Time    `stbl:"datetime_nn"`
	DatetimeNnUnique            time.Time    `stbl:"datetime_nn_unique,PRIMARY_KEY"`
	DatetimeNnCheck             time.Time    `stbl:"datetime_nn_check"`
	DatetimeNnRef               time.Time    `stbl:"datetime_nn_ref"`
	DatetimeNnDefConst          time.Time    `stbl:"datetime_nn_def_const"`
	DatetimeNnDefFunc           time.Time    `stbl:"datetime_nn_def_func"`
	DatetimeNnUniqueCheck       time.Time    `stbl:"datetime_nn_unique_check"`
	DatetimeUnique              sql.NullTime `stbl:"datetime_unique"`
	DatetimeUniqueCheck         sql.NullTime `stbl:"datetime_unique_check"`
	DatetimeUniqueRef           sql.NullTime `stbl:"datetime_unique_ref"`
	DatetimeUniqueDefConst      sql.NullTime `stbl:"datetime_unique_def_const"`
	DatetimeUniqueDefFunc       sql.NullTime `stbl:"datetime_unique_def_func"`
	DatetimeCheck               sql.NullTime `stbl:"datetime_check"`
	DatetimeCheckRef            sql.NullTime `stbl:"datetime_check_ref"`
	DatetimeCheckDefConst       sql.NullTime `stbl:"datetime_check_def_const"`
	DatetimeCheckDefFunc        sql.NullTime `stbl:"datetime_check_def_func"`
	DatetimeRef                 sql.NullTime `stbl:"datetime_ref"`
	DatetimeRefUniqueCheck      sql.NullTime `stbl:"datetime_ref_unique_check"`
	DatetimeDefConst            sql.NullTime `stbl:"datetime_def_const"`
	DatetimeDefConstUniqueCheck sql.NullTime `stbl:"datetime_def_const_unique_check"`
	DatetimeDefFunc             sql.NullTime `stbl:"datetime_def_func"`
	DatetimeDefFuncUniqueCheck  sql.NullTime `stbl:"datetime_def_func_unique_check"`
}
