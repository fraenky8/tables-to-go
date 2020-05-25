package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type TimeWithoutTimeZone struct {
	TimeWithoutTimeZone                    pg.NullTime `db:"time_without_time_zone"`
	TimeWithoutTimeZoneNn                  time.Time   `db:"time_without_time_zone_nn"`
	TimeWithoutTimeZoneNnUnique            time.Time   `db:"time_without_time_zone_nn_unique"`
	TimeWithoutTimeZoneNnCheck             time.Time   `db:"time_without_time_zone_nn_check"`
	TimeWithoutTimeZoneNnRef               time.Time   `db:"time_without_time_zone_nn_ref"`
	TimeWithoutTimeZoneNnDefConst          time.Time   `db:"time_without_time_zone_nn_def_const"`
	TimeWithoutTimeZoneNnDefFunc           time.Time   `db:"time_without_time_zone_nn_def_func"`
	TimeWithoutTimeZoneNnUniqueCheck       time.Time   `db:"time_without_time_zone_nn_unique_check"`
	TimeWithoutTimeZoneUnique              pg.NullTime `db:"time_without_time_zone_unique"`
	TimeWithoutTimeZoneUniqueCheck         pg.NullTime `db:"time_without_time_zone_unique_check"`
	TimeWithoutTimeZoneUniqueRef           pg.NullTime `db:"time_without_time_zone_unique_ref"`
	TimeWithoutTimeZoneUniqueDefConst      pg.NullTime `db:"time_without_time_zone_unique_def_const"`
	TimeWithoutTimeZoneUniqueDefFunc       pg.NullTime `db:"time_without_time_zone_unique_def_func"`
	TimeWithoutTimeZoneCheck               pg.NullTime `db:"time_without_time_zone_check"`
	TimeWithoutTimeZoneCheckRef            pg.NullTime `db:"time_without_time_zone_check_ref"`
	TimeWithoutTimeZoneCheckDefConst       pg.NullTime `db:"time_without_time_zone_check_def_const"`
	TimeWithoutTimeZoneCheckDefFunc        pg.NullTime `db:"time_without_time_zone_check_def_func"`
	TimeWithoutTimeZoneRef                 pg.NullTime `db:"time_without_time_zone_ref"`
	TimeWithoutTimeZoneRefDefConst         pg.NullTime `db:"time_without_time_zone_ref_def_const"`
	TimeWithoutTimeZoneRefDefFunc          pg.NullTime `db:"time_without_time_zone_ref_def_func"`
	TimeWithoutTimeZoneRefUniqueCheck      pg.NullTime `db:"time_without_time_zone_ref_unique_check"`
	TimeWithoutTimeZoneDefConst            pg.NullTime `db:"time_without_time_zone_def_const"`
	TimeWithoutTimeZoneDefConstUniqueCheck pg.NullTime `db:"time_without_time_zone_def_const_unique_check"`
	TimeWithoutTimeZoneDefFunc             pg.NullTime `db:"time_without_time_zone_def_func"`
	TimeWithoutTimeZoneDefFuncUniqueCheck  pg.NullTime `db:"time_without_time_zone_def_func_unique_check"`
}
