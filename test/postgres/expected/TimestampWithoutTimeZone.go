package dto

import (
	"database/sql"
	"time"
)

type TimestampWithoutTimeZone struct {
	TimestampWithoutTimeZone                    sql.NullTime `db:"timestamp_without_time_zone"`
	TimestampWithoutTimeZoneNn                  time.Time    `db:"timestamp_without_time_zone_nn"`
	TimestampWithoutTimeZoneNnUnique            time.Time    `db:"timestamp_without_time_zone_nn_unique"`
	TimestampWithoutTimeZoneNnCheck             time.Time    `db:"timestamp_without_time_zone_nn_check"`
	TimestampWithoutTimeZoneNnRef               time.Time    `db:"timestamp_without_time_zone_nn_ref"`
	TimestampWithoutTimeZoneNnDefConst          time.Time    `db:"timestamp_without_time_zone_nn_def_const"`
	TimestampWithoutTimeZoneNnDefFunc           time.Time    `db:"timestamp_without_time_zone_nn_def_func"`
	TimestampWithoutTimeZoneNnUniqueCheck       time.Time    `db:"timestamp_without_time_zone_nn_unique_check"`
	TimestampWithoutTimeZoneUnique              sql.NullTime `db:"timestamp_without_time_zone_unique"`
	TimestampWithoutTimeZoneUniqueCheck         sql.NullTime `db:"timestamp_without_time_zone_unique_check"`
	TimestampWithoutTimeZoneUniqueRef           sql.NullTime `db:"timestamp_without_time_zone_unique_ref"`
	TimestampWithoutTimeZoneUniqueDefConst      sql.NullTime `db:"timestamp_without_time_zone_unique_def_const"`
	TimestampWithoutTimeZoneUniqueDefFunc       sql.NullTime `db:"timestamp_without_time_zone_unique_def_func"`
	TimestampWithoutTimeZoneCheck               sql.NullTime `db:"timestamp_without_time_zone_check"`
	TimestampWithoutTimeZoneCheckRef            sql.NullTime `db:"timestamp_without_time_zone_check_ref"`
	TimestampWithoutTimeZoneCheckDefConst       sql.NullTime `db:"timestamp_without_time_zone_check_def_const"`
	TimestampWithoutTimeZoneCheckDefFunc        sql.NullTime `db:"timestamp_without_time_zone_check_def_func"`
	TimestampWithoutTimeZoneRef                 sql.NullTime `db:"timestamp_without_time_zone_ref"`
	TimestampWithoutTimeZoneRefDefConst         sql.NullTime `db:"timestamp_without_time_zone_ref_def_const"`
	TimestampWithoutTimeZoneRefDefFunc          sql.NullTime `db:"timestamp_without_time_zone_ref_def_func"`
	TimestampWithoutTimeZoneRefUniqueCheck      sql.NullTime `db:"timestamp_without_time_zone_ref_unique_check"`
	TimestampWithoutTimeZoneDefConst            sql.NullTime `db:"timestamp_without_time_zone_def_const"`
	TimestampWithoutTimeZoneDefConstUniqueCheck sql.NullTime `db:"timestamp_without_time_zone_def_const_unique_check"`
	TimestampWithoutTimeZoneDefFunc             sql.NullTime `db:"timestamp_without_time_zone_def_func"`
	TimestampWithoutTimeZoneDefFuncUniqueCheck  sql.NullTime `db:"timestamp_without_time_zone_def_func_unique_check"`
}
