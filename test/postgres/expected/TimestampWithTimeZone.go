package dto

import (
	"database/sql"
	"time"
)

type TimestampWithTimeZone struct {
	TimestampWithTimeZone                    sql.NullTime `db:"timestamp_with_time_zone"`
	TimestampWithTimeZoneNn                  time.Time    `db:"timestamp_with_time_zone_nn"`
	TimestampWithTimeZoneNnUnique            time.Time    `db:"timestamp_with_time_zone_nn_unique"`
	TimestampWithTimeZoneNnCheck             time.Time    `db:"timestamp_with_time_zone_nn_check"`
	TimestampWithTimeZoneNnRef               time.Time    `db:"timestamp_with_time_zone_nn_ref"`
	TimestampWithTimeZoneNnDefConst          time.Time    `db:"timestamp_with_time_zone_nn_def_const"`
	TimestampWithTimeZoneNnDefFunc           time.Time    `db:"timestamp_with_time_zone_nn_def_func"`
	TimestampWithTimeZoneNnUniqueCheck       time.Time    `db:"timestamp_with_time_zone_nn_unique_check"`
	TimestampWithTimeZoneUnique              sql.NullTime `db:"timestamp_with_time_zone_unique"`
	TimestampWithTimeZoneUniqueCheck         sql.NullTime `db:"timestamp_with_time_zone_unique_check"`
	TimestampWithTimeZoneUniqueRef           sql.NullTime `db:"timestamp_with_time_zone_unique_ref"`
	TimestampWithTimeZoneUniqueDefConst      sql.NullTime `db:"timestamp_with_time_zone_unique_def_const"`
	TimestampWithTimeZoneUniqueDefFunc       sql.NullTime `db:"timestamp_with_time_zone_unique_def_func"`
	TimestampWithTimeZoneCheck               sql.NullTime `db:"timestamp_with_time_zone_check"`
	TimestampWithTimeZoneCheckRef            sql.NullTime `db:"timestamp_with_time_zone_check_ref"`
	TimestampWithTimeZoneCheckDefConst       sql.NullTime `db:"timestamp_with_time_zone_check_def_const"`
	TimestampWithTimeZoneCheckDefFunc        sql.NullTime `db:"timestamp_with_time_zone_check_def_func"`
	TimestampWithTimeZoneRef                 sql.NullTime `db:"timestamp_with_time_zone_ref"`
	TimestampWithTimeZoneRefDefConst         sql.NullTime `db:"timestamp_with_time_zone_ref_def_const"`
	TimestampWithTimeZoneRefDefFunc          sql.NullTime `db:"timestamp_with_time_zone_ref_def_func"`
	TimestampWithTimeZoneRefUniqueCheck      sql.NullTime `db:"timestamp_with_time_zone_ref_unique_check"`
	TimestampWithTimeZoneDefConst            sql.NullTime `db:"timestamp_with_time_zone_def_const"`
	TimestampWithTimeZoneDefConstUniqueCheck sql.NullTime `db:"timestamp_with_time_zone_def_const_unique_check"`
	TimestampWithTimeZoneDefFunc             sql.NullTime `db:"timestamp_with_time_zone_def_func"`
	TimestampWithTimeZoneDefFuncUniqueCheck  sql.NullTime `db:"timestamp_with_time_zone_def_func_unique_check"`
}
