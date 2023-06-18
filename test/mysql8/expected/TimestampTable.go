package dto

import (
	"database/sql"
	"time"
)

type TimestampTable struct {
	Timestamp                    sql.NullTime `db:"timestamp"`
	TimestampNn                  time.Time    `db:"timestamp_nn"`
	TimestampNnUnique            time.Time    `db:"timestamp_nn_unique"`
	TimestampNnCheck             time.Time    `db:"timestamp_nn_check"`
	TimestampNnRef               time.Time    `db:"timestamp_nn_ref"`
	TimestampNnDefConst          time.Time    `db:"timestamp_nn_def_const"`
	TimestampNnDefFunc           time.Time    `db:"timestamp_nn_def_func"`
	TimestampNnUniqueCheck       time.Time    `db:"timestamp_nn_unique_check"`
	TimestampUnique              sql.NullTime `db:"timestamp_unique"`
	TimestampUniqueCheck         sql.NullTime `db:"timestamp_unique_check"`
	TimestampUniqueRef           sql.NullTime `db:"timestamp_unique_ref"`
	TimestampUniqueDefConst      sql.NullTime `db:"timestamp_unique_def_const"`
	TimestampUniqueDefFunc       sql.NullTime `db:"timestamp_unique_def_func"`
	TimestampCheck               sql.NullTime `db:"timestamp_check"`
	TimestampCheckRef            sql.NullTime `db:"timestamp_check_ref"`
	TimestampCheckDefConst       sql.NullTime `db:"timestamp_check_def_const"`
	TimestampCheckDefFunc        sql.NullTime `db:"timestamp_check_def_func"`
	TimestampRef                 sql.NullTime `db:"timestamp_ref"`
	TimestampRefUniqueCheck      sql.NullTime `db:"timestamp_ref_unique_check"`
	TimestampDefConst            sql.NullTime `db:"timestamp_def_const"`
	TimestampDefConstUniqueCheck sql.NullTime `db:"timestamp_def_const_unique_check"`
	TimestampDefFunc             sql.NullTime `db:"timestamp_def_func"`
	TimestampDefFuncUniqueCheck  sql.NullTime `db:"timestamp_def_func_unique_check"`
}
