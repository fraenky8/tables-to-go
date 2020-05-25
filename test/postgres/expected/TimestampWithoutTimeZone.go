package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type TimestampWithoutTimeZone struct {
	TimestampWithoutTimeZone                    pg.NullTime `db:"timestamp_without_time_zone"`
	TimestampWithoutTimeZoneNn                  time.Time   `db:"timestamp_without_time_zone_nn"`
	TimestampWithoutTimeZoneNnUnique            time.Time   `db:"timestamp_without_time_zone_nn_unique"`
	TimestampWithoutTimeZoneNnCheck             time.Time   `db:"timestamp_without_time_zone_nn_check"`
	TimestampWithoutTimeZoneNnRef               time.Time   `db:"timestamp_without_time_zone_nn_ref"`
	TimestampWithoutTimeZoneNnDefConst          time.Time   `db:"timestamp_without_time_zone_nn_def_const"`
	TimestampWithoutTimeZoneNnDefFunc           time.Time   `db:"timestamp_without_time_zone_nn_def_func"`
	TimestampWithoutTimeZoneNnUniqueCheck       time.Time   `db:"timestamp_without_time_zone_nn_unique_check"`
	TimestampWithoutTimeZoneUnique              pg.NullTime `db:"timestamp_without_time_zone_unique"`
	TimestampWithoutTimeZoneUniqueCheck         pg.NullTime `db:"timestamp_without_time_zone_unique_check"`
	TimestampWithoutTimeZoneUniqueRef           pg.NullTime `db:"timestamp_without_time_zone_unique_ref"`
	TimestampWithoutTimeZoneUniqueDefConst      pg.NullTime `db:"timestamp_without_time_zone_unique_def_const"`
	TimestampWithoutTimeZoneUniqueDefFunc       pg.NullTime `db:"timestamp_without_time_zone_unique_def_func"`
	TimestampWithoutTimeZoneCheck               pg.NullTime `db:"timestamp_without_time_zone_check"`
	TimestampWithoutTimeZoneCheckRef            pg.NullTime `db:"timestamp_without_time_zone_check_ref"`
	TimestampWithoutTimeZoneCheckDefConst       pg.NullTime `db:"timestamp_without_time_zone_check_def_const"`
	TimestampWithoutTimeZoneCheckDefFunc        pg.NullTime `db:"timestamp_without_time_zone_check_def_func"`
	TimestampWithoutTimeZoneRef                 pg.NullTime `db:"timestamp_without_time_zone_ref"`
	TimestampWithoutTimeZoneRefDefConst         pg.NullTime `db:"timestamp_without_time_zone_ref_def_const"`
	TimestampWithoutTimeZoneRefDefFunc          pg.NullTime `db:"timestamp_without_time_zone_ref_def_func"`
	TimestampWithoutTimeZoneRefUniqueCheck      pg.NullTime `db:"timestamp_without_time_zone_ref_unique_check"`
	TimestampWithoutTimeZoneDefConst            pg.NullTime `db:"timestamp_without_time_zone_def_const"`
	TimestampWithoutTimeZoneDefConstUniqueCheck pg.NullTime `db:"timestamp_without_time_zone_def_const_unique_check"`
	TimestampWithoutTimeZoneDefFunc             pg.NullTime `db:"timestamp_without_time_zone_def_func"`
	TimestampWithoutTimeZoneDefFuncUniqueCheck  pg.NullTime `db:"timestamp_without_time_zone_def_func_unique_check"`
}
