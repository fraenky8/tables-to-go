package dto

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type TimestampTable struct {
	Timestamp                    mysql.NullTime `db:"timestamp"`
	TimestampNn                  time.Time      `db:"timestamp_nn"`
	TimestampNnUnique            time.Time      `db:"timestamp_nn_unique"`
	TimestampNnCheck             time.Time      `db:"timestamp_nn_check"`
	TimestampNnRef               time.Time      `db:"timestamp_nn_ref"`
	TimestampNnDefConst          time.Time      `db:"timestamp_nn_def_const"`
	TimestampNnDefFunc           time.Time      `db:"timestamp_nn_def_func"`
	TimestampNnUniqueCheck       time.Time      `db:"timestamp_nn_unique_check"`
	TimestampUnique              mysql.NullTime `db:"timestamp_unique"`
	TimestampUniqueCheck         mysql.NullTime `db:"timestamp_unique_check"`
	TimestampUniqueRef           mysql.NullTime `db:"timestamp_unique_ref"`
	TimestampUniqueDefConst      mysql.NullTime `db:"timestamp_unique_def_const"`
	TimestampUniqueDefFunc       mysql.NullTime `db:"timestamp_unique_def_func"`
	TimestampCheck               mysql.NullTime `db:"timestamp_check"`
	TimestampCheckRef            mysql.NullTime `db:"timestamp_check_ref"`
	TimestampCheckDefConst       mysql.NullTime `db:"timestamp_check_def_const"`
	TimestampCheckDefFunc        mysql.NullTime `db:"timestamp_check_def_func"`
	TimestampRef                 mysql.NullTime `db:"timestamp_ref"`
	TimestampRefUniqueCheck      mysql.NullTime `db:"timestamp_ref_unique_check"`
	TimestampDefConst            mysql.NullTime `db:"timestamp_def_const"`
	TimestampDefConstUniqueCheck mysql.NullTime `db:"timestamp_def_const_unique_check"`
	TimestampDefFunc             mysql.NullTime `db:"timestamp_def_func"`
	TimestampDefFuncUniqueCheck  mysql.NullTime `db:"timestamp_def_func_unique_check"`
}
