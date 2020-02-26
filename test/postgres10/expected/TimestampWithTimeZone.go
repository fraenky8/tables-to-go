package dto

import (
	"time"

	pg "github.com/lib/pq"
)

type TimestampWithTimeZone struct {
	TimestampWithTimeZone                    pg.NullTime `db:"timestamp_with_time_zone"`
	TimestampWithTimeZoneNn                  time.Time   `db:"timestamp_with_time_zone_nn"`
	TimestampWithTimeZoneNnUnique            time.Time   `db:"timestamp_with_time_zone_nn_unique"`
	TimestampWithTimeZoneNnCheck             time.Time   `db:"timestamp_with_time_zone_nn_check"`
	TimestampWithTimeZoneNnRef               time.Time   `db:"timestamp_with_time_zone_nn_ref"`
	TimestampWithTimeZoneNnDefConst          time.Time   `db:"timestamp_with_time_zone_nn_def_const"`
	TimestampWithTimeZoneNnDefFunc           time.Time   `db:"timestamp_with_time_zone_nn_def_func"`
	TimestampWithTimeZoneNnUniqueCheck       time.Time   `db:"timestamp_with_time_zone_nn_unique_check"`
	TimestampWithTimeZoneUnique              pg.NullTime `db:"timestamp_with_time_zone_unique"`
	TimestampWithTimeZoneUniqueCheck         pg.NullTime `db:"timestamp_with_time_zone_unique_check"`
	TimestampWithTimeZoneUniqueRef           pg.NullTime `db:"timestamp_with_time_zone_unique_ref"`
	TimestampWithTimeZoneUniqueDefConst      pg.NullTime `db:"timestamp_with_time_zone_unique_def_const"`
	TimestampWithTimeZoneUniqueDefFunc       pg.NullTime `db:"timestamp_with_time_zone_unique_def_func"`
	TimestampWithTimeZoneCheck               pg.NullTime `db:"timestamp_with_time_zone_check"`
	TimestampWithTimeZoneCheckRef            pg.NullTime `db:"timestamp_with_time_zone_check_ref"`
	TimestampWithTimeZoneCheckDefConst       pg.NullTime `db:"timestamp_with_time_zone_check_def_const"`
	TimestampWithTimeZoneCheckDefFunc        pg.NullTime `db:"timestamp_with_time_zone_check_def_func"`
	TimestampWithTimeZoneRef                 pg.NullTime `db:"timestamp_with_time_zone_ref"`
	TimestampWithTimeZoneRefDefConst         pg.NullTime `db:"timestamp_with_time_zone_ref_def_const"`
	TimestampWithTimeZoneRefDefFunc          pg.NullTime `db:"timestamp_with_time_zone_ref_def_func"`
	TimestampWithTimeZoneRefUniqueCheck      pg.NullTime `db:"timestamp_with_time_zone_ref_unique_check"`
	TimestampWithTimeZoneDefConst            pg.NullTime `db:"timestamp_with_time_zone_def_const"`
	TimestampWithTimeZoneDefConstUniqueCheck pg.NullTime `db:"timestamp_with_time_zone_def_const_unique_check"`
	TimestampWithTimeZoneDefFunc             pg.NullTime `db:"timestamp_with_time_zone_def_func"`
	TimestampWithTimeZoneDefFuncUniqueCheck  pg.NullTime `db:"timestamp_with_time_zone_def_func_unique_check"`
}
