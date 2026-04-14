package dto

import (
	"time"
)

type TimeWithTimeZone struct {
	TimeWithTimeZone                    *time.Time `db:"time_with_time_zone"`
	TimeWithTimeZoneNn                  time.Time  `db:"time_with_time_zone_nn"`
	TimeWithTimeZoneNnUnique            time.Time  `db:"time_with_time_zone_nn_unique"`
	TimeWithTimeZoneNnCheck             time.Time  `db:"time_with_time_zone_nn_check"`
	TimeWithTimeZoneNnRef               time.Time  `db:"time_with_time_zone_nn_ref"`
	TimeWithTimeZoneNnDefConst          time.Time  `db:"time_with_time_zone_nn_def_const"`
	TimeWithTimeZoneNnDefFunc           time.Time  `db:"time_with_time_zone_nn_def_func"`
	TimeWithTimeZoneNnUniqueCheck       time.Time  `db:"time_with_time_zone_nn_unique_check"`
	TimeWithTimeZoneUnique              *time.Time `db:"time_with_time_zone_unique"`
	TimeWithTimeZoneUniqueCheck         *time.Time `db:"time_with_time_zone_unique_check"`
	TimeWithTimeZoneUniqueRef           *time.Time `db:"time_with_time_zone_unique_ref"`
	TimeWithTimeZoneUniqueDefConst      *time.Time `db:"time_with_time_zone_unique_def_const"`
	TimeWithTimeZoneUniqueDefFunc       *time.Time `db:"time_with_time_zone_unique_def_func"`
	TimeWithTimeZoneCheck               *time.Time `db:"time_with_time_zone_check"`
	TimeWithTimeZoneCheckRef            *time.Time `db:"time_with_time_zone_check_ref"`
	TimeWithTimeZoneCheckDefConst       *time.Time `db:"time_with_time_zone_check_def_const"`
	TimeWithTimeZoneCheckDefFunc        *time.Time `db:"time_with_time_zone_check_def_func"`
	TimeWithTimeZoneRef                 *time.Time `db:"time_with_time_zone_ref"`
	TimeWithTimeZoneRefDefConst         *time.Time `db:"time_with_time_zone_ref_def_const"`
	TimeWithTimeZoneRefDefFunc          *time.Time `db:"time_with_time_zone_ref_def_func"`
	TimeWithTimeZoneRefUniqueCheck      *time.Time `db:"time_with_time_zone_ref_unique_check"`
	TimeWithTimeZoneDefConst            *time.Time `db:"time_with_time_zone_def_const"`
	TimeWithTimeZoneDefConstUniqueCheck *time.Time `db:"time_with_time_zone_def_const_unique_check"`
	TimeWithTimeZoneDefFunc             *time.Time `db:"time_with_time_zone_def_func"`
	TimeWithTimeZoneDefFuncUniqueCheck  *time.Time `db:"time_with_time_zone_def_func_unique_check"`
}
