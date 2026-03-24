package dto

import (
	"time"
)

type TimestampWithoutTimeZone struct {
	TimestampWithoutTimeZone                    *time.Time `db:"timestamp_without_time_zone"`
	TimestampWithoutTimeZoneNn                  time.Time  `db:"timestamp_without_time_zone_nn"`
	TimestampWithoutTimeZoneNnUnique            time.Time  `db:"timestamp_without_time_zone_nn_unique"`
	TimestampWithoutTimeZoneNnCheck             time.Time  `db:"timestamp_without_time_zone_nn_check"`
	TimestampWithoutTimeZoneNnRef               time.Time  `db:"timestamp_without_time_zone_nn_ref"`
	TimestampWithoutTimeZoneNnDefConst          time.Time  `db:"timestamp_without_time_zone_nn_def_const"`
	TimestampWithoutTimeZoneNnDefFunc           time.Time  `db:"timestamp_without_time_zone_nn_def_func"`
	TimestampWithoutTimeZoneNnUniqueCheck       time.Time  `db:"timestamp_without_time_zone_nn_unique_check"`
	TimestampWithoutTimeZoneUnique              *time.Time `db:"timestamp_without_time_zone_unique"`
	TimestampWithoutTimeZoneUniqueCheck         *time.Time `db:"timestamp_without_time_zone_unique_check"`
	TimestampWithoutTimeZoneUniqueRef           *time.Time `db:"timestamp_without_time_zone_unique_ref"`
	TimestampWithoutTimeZoneUniqueDefConst      *time.Time `db:"timestamp_without_time_zone_unique_def_const"`
	TimestampWithoutTimeZoneUniqueDefFunc       *time.Time `db:"timestamp_without_time_zone_unique_def_func"`
	TimestampWithoutTimeZoneCheck               *time.Time `db:"timestamp_without_time_zone_check"`
	TimestampWithoutTimeZoneCheckRef            *time.Time `db:"timestamp_without_time_zone_check_ref"`
	TimestampWithoutTimeZoneCheckDefConst       *time.Time `db:"timestamp_without_time_zone_check_def_const"`
	TimestampWithoutTimeZoneCheckDefFunc        *time.Time `db:"timestamp_without_time_zone_check_def_func"`
	TimestampWithoutTimeZoneRef                 *time.Time `db:"timestamp_without_time_zone_ref"`
	TimestampWithoutTimeZoneRefDefConst         *time.Time `db:"timestamp_without_time_zone_ref_def_const"`
	TimestampWithoutTimeZoneRefDefFunc          *time.Time `db:"timestamp_without_time_zone_ref_def_func"`
	TimestampWithoutTimeZoneRefUniqueCheck      *time.Time `db:"timestamp_without_time_zone_ref_unique_check"`
	TimestampWithoutTimeZoneDefConst            *time.Time `db:"timestamp_without_time_zone_def_const"`
	TimestampWithoutTimeZoneDefConstUniqueCheck *time.Time `db:"timestamp_without_time_zone_def_const_unique_check"`
	TimestampWithoutTimeZoneDefFunc             *time.Time `db:"timestamp_without_time_zone_def_func"`
	TimestampWithoutTimeZoneDefFuncUniqueCheck  *time.Time `db:"timestamp_without_time_zone_def_func_unique_check"`
}
