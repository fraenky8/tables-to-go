package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type TimeWithTimeZone struct {
	TimeWithTimeZone                    pg.NullTime `db:"time_with_time_zone"`
	TimeWithTimeZoneNn                  time.Time   `db:"time_with_time_zone_nn"`
	TimeWithTimeZoneNnUnique            time.Time   `db:"time_with_time_zone_nn_unique"`
	TimeWithTimeZoneNnCheck             time.Time   `db:"time_with_time_zone_nn_check"`
	TimeWithTimeZoneNnRef               time.Time   `db:"time_with_time_zone_nn_ref"`
	TimeWithTimeZoneNnDefConst          time.Time   `db:"time_with_time_zone_nn_def_const"`
	TimeWithTimeZoneNnDefFunc           time.Time   `db:"time_with_time_zone_nn_def_func"`
	TimeWithTimeZoneNnUniqueCheck       time.Time   `db:"time_with_time_zone_nn_unique_check"`
	TimeWithTimeZoneUnique              pg.NullTime `db:"time_with_time_zone_unique"`
	TimeWithTimeZoneUniqueCheck         pg.NullTime `db:"time_with_time_zone_unique_check"`
	TimeWithTimeZoneUniqueRef           pg.NullTime `db:"time_with_time_zone_unique_ref"`
	TimeWithTimeZoneUniqueDefConst      pg.NullTime `db:"time_with_time_zone_unique_def_const"`
	TimeWithTimeZoneUniqueDefFunc       pg.NullTime `db:"time_with_time_zone_unique_def_func"`
	TimeWithTimeZoneCheck               pg.NullTime `db:"time_with_time_zone_check"`
	TimeWithTimeZoneCheckRef            pg.NullTime `db:"time_with_time_zone_check_ref"`
	TimeWithTimeZoneCheckDefConst       pg.NullTime `db:"time_with_time_zone_check_def_const"`
	TimeWithTimeZoneCheckDefFunc        pg.NullTime `db:"time_with_time_zone_check_def_func"`
	TimeWithTimeZoneRef                 pg.NullTime `db:"time_with_time_zone_ref"`
	TimeWithTimeZoneRefDefConst         pg.NullTime `db:"time_with_time_zone_ref_def_const"`
	TimeWithTimeZoneRefDefFunc          pg.NullTime `db:"time_with_time_zone_ref_def_func"`
	TimeWithTimeZoneRefUniqueCheck      pg.NullTime `db:"time_with_time_zone_ref_unique_check"`
	TimeWithTimeZoneDefConst            pg.NullTime `db:"time_with_time_zone_def_const"`
	TimeWithTimeZoneDefConstUniqueCheck pg.NullTime `db:"time_with_time_zone_def_const_unique_check"`
	TimeWithTimeZoneDefFunc             pg.NullTime `db:"time_with_time_zone_def_func"`
	TimeWithTimeZoneDefFuncUniqueCheck  pg.NullTime `db:"time_with_time_zone_def_func_unique_check"`
}
