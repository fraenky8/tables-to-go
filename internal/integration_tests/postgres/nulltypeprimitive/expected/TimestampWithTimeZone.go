package dto

import (
	"time"
)

type TimestampWithTimeZone struct {
	TimestampWithTimeZone                    *time.Time `db:"timestamp_with_time_zone"`
	TimestampWithTimeZoneNn                  time.Time  `db:"timestamp_with_time_zone_nn"`
	TimestampWithTimeZoneNnUnique            time.Time  `db:"timestamp_with_time_zone_nn_unique"`
	TimestampWithTimeZoneNnCheck             time.Time  `db:"timestamp_with_time_zone_nn_check"`
	TimestampWithTimeZoneNnRef               time.Time  `db:"timestamp_with_time_zone_nn_ref"`
	TimestampWithTimeZoneNnDefConst          time.Time  `db:"timestamp_with_time_zone_nn_def_const"`
	TimestampWithTimeZoneNnDefFunc           time.Time  `db:"timestamp_with_time_zone_nn_def_func"`
	TimestampWithTimeZoneNnUniqueCheck       time.Time  `db:"timestamp_with_time_zone_nn_unique_check"`
	TimestampWithTimeZoneUnique              *time.Time `db:"timestamp_with_time_zone_unique"`
	TimestampWithTimeZoneUniqueCheck         *time.Time `db:"timestamp_with_time_zone_unique_check"`
	TimestampWithTimeZoneUniqueRef           *time.Time `db:"timestamp_with_time_zone_unique_ref"`
	TimestampWithTimeZoneUniqueDefConst      *time.Time `db:"timestamp_with_time_zone_unique_def_const"`
	TimestampWithTimeZoneUniqueDefFunc       *time.Time `db:"timestamp_with_time_zone_unique_def_func"`
	TimestampWithTimeZoneCheck               *time.Time `db:"timestamp_with_time_zone_check"`
	TimestampWithTimeZoneCheckRef            *time.Time `db:"timestamp_with_time_zone_check_ref"`
	TimestampWithTimeZoneCheckDefConst       *time.Time `db:"timestamp_with_time_zone_check_def_const"`
	TimestampWithTimeZoneCheckDefFunc        *time.Time `db:"timestamp_with_time_zone_check_def_func"`
	TimestampWithTimeZoneRef                 *time.Time `db:"timestamp_with_time_zone_ref"`
	TimestampWithTimeZoneRefDefConst         *time.Time `db:"timestamp_with_time_zone_ref_def_const"`
	TimestampWithTimeZoneRefDefFunc          *time.Time `db:"timestamp_with_time_zone_ref_def_func"`
	TimestampWithTimeZoneRefUniqueCheck      *time.Time `db:"timestamp_with_time_zone_ref_unique_check"`
	TimestampWithTimeZoneDefConst            *time.Time `db:"timestamp_with_time_zone_def_const"`
	TimestampWithTimeZoneDefConstUniqueCheck *time.Time `db:"timestamp_with_time_zone_def_const_unique_check"`
	TimestampWithTimeZoneDefFunc             *time.Time `db:"timestamp_with_time_zone_def_func"`
	TimestampWithTimeZoneDefFuncUniqueCheck  *time.Time `db:"timestamp_with_time_zone_def_func_unique_check"`
}
