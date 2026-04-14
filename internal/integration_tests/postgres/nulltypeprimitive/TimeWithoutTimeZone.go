package dto

import (
	"time"
)

type TimeWithoutTimeZone struct {
	TimeWithoutTimeZone                    *time.Time `db:"time_without_time_zone"`
	TimeWithoutTimeZoneNn                  time.Time  `db:"time_without_time_zone_nn"`
	TimeWithoutTimeZoneNnUnique            time.Time  `db:"time_without_time_zone_nn_unique"`
	TimeWithoutTimeZoneNnCheck             time.Time  `db:"time_without_time_zone_nn_check"`
	TimeWithoutTimeZoneNnRef               time.Time  `db:"time_without_time_zone_nn_ref"`
	TimeWithoutTimeZoneNnDefConst          time.Time  `db:"time_without_time_zone_nn_def_const"`
	TimeWithoutTimeZoneNnDefFunc           time.Time  `db:"time_without_time_zone_nn_def_func"`
	TimeWithoutTimeZoneNnUniqueCheck       time.Time  `db:"time_without_time_zone_nn_unique_check"`
	TimeWithoutTimeZoneUnique              *time.Time `db:"time_without_time_zone_unique"`
	TimeWithoutTimeZoneUniqueCheck         *time.Time `db:"time_without_time_zone_unique_check"`
	TimeWithoutTimeZoneUniqueRef           *time.Time `db:"time_without_time_zone_unique_ref"`
	TimeWithoutTimeZoneUniqueDefConst      *time.Time `db:"time_without_time_zone_unique_def_const"`
	TimeWithoutTimeZoneUniqueDefFunc       *time.Time `db:"time_without_time_zone_unique_def_func"`
	TimeWithoutTimeZoneCheck               *time.Time `db:"time_without_time_zone_check"`
	TimeWithoutTimeZoneCheckRef            *time.Time `db:"time_without_time_zone_check_ref"`
	TimeWithoutTimeZoneCheckDefConst       *time.Time `db:"time_without_time_zone_check_def_const"`
	TimeWithoutTimeZoneCheckDefFunc        *time.Time `db:"time_without_time_zone_check_def_func"`
	TimeWithoutTimeZoneRef                 *time.Time `db:"time_without_time_zone_ref"`
	TimeWithoutTimeZoneRefDefConst         *time.Time `db:"time_without_time_zone_ref_def_const"`
	TimeWithoutTimeZoneRefDefFunc          *time.Time `db:"time_without_time_zone_ref_def_func"`
	TimeWithoutTimeZoneRefUniqueCheck      *time.Time `db:"time_without_time_zone_ref_unique_check"`
	TimeWithoutTimeZoneDefConst            *time.Time `db:"time_without_time_zone_def_const"`
	TimeWithoutTimeZoneDefConstUniqueCheck *time.Time `db:"time_without_time_zone_def_const_unique_check"`
	TimeWithoutTimeZoneDefFunc             *time.Time `db:"time_without_time_zone_def_func"`
	TimeWithoutTimeZoneDefFuncUniqueCheck  *time.Time `db:"time_without_time_zone_def_func_unique_check"`
}
