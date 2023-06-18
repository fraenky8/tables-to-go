package dto

import (
	"database/sql"
	"time"
)

type TimeWithTimeZone struct {
	TimeWithTimeZone                    sql.NullTime `db:"time_with_time_zone"`
	TimeWithTimeZoneNn                  time.Time    `db:"time_with_time_zone_nn"`
	TimeWithTimeZoneNnUnique            time.Time    `db:"time_with_time_zone_nn_unique"`
	TimeWithTimeZoneNnCheck             time.Time    `db:"time_with_time_zone_nn_check"`
	TimeWithTimeZoneNnRef               time.Time    `db:"time_with_time_zone_nn_ref"`
	TimeWithTimeZoneNnDefConst          time.Time    `db:"time_with_time_zone_nn_def_const"`
	TimeWithTimeZoneNnDefFunc           time.Time    `db:"time_with_time_zone_nn_def_func"`
	TimeWithTimeZoneNnUniqueCheck       time.Time    `db:"time_with_time_zone_nn_unique_check"`
	TimeWithTimeZoneUnique              sql.NullTime `db:"time_with_time_zone_unique"`
	TimeWithTimeZoneUniqueCheck         sql.NullTime `db:"time_with_time_zone_unique_check"`
	TimeWithTimeZoneUniqueRef           sql.NullTime `db:"time_with_time_zone_unique_ref"`
	TimeWithTimeZoneUniqueDefConst      sql.NullTime `db:"time_with_time_zone_unique_def_const"`
	TimeWithTimeZoneUniqueDefFunc       sql.NullTime `db:"time_with_time_zone_unique_def_func"`
	TimeWithTimeZoneCheck               sql.NullTime `db:"time_with_time_zone_check"`
	TimeWithTimeZoneCheckRef            sql.NullTime `db:"time_with_time_zone_check_ref"`
	TimeWithTimeZoneCheckDefConst       sql.NullTime `db:"time_with_time_zone_check_def_const"`
	TimeWithTimeZoneCheckDefFunc        sql.NullTime `db:"time_with_time_zone_check_def_func"`
	TimeWithTimeZoneRef                 sql.NullTime `db:"time_with_time_zone_ref"`
	TimeWithTimeZoneRefDefConst         sql.NullTime `db:"time_with_time_zone_ref_def_const"`
	TimeWithTimeZoneRefDefFunc          sql.NullTime `db:"time_with_time_zone_ref_def_func"`
	TimeWithTimeZoneRefUniqueCheck      sql.NullTime `db:"time_with_time_zone_ref_unique_check"`
	TimeWithTimeZoneDefConst            sql.NullTime `db:"time_with_time_zone_def_const"`
	TimeWithTimeZoneDefConstUniqueCheck sql.NullTime `db:"time_with_time_zone_def_const_unique_check"`
	TimeWithTimeZoneDefFunc             sql.NullTime `db:"time_with_time_zone_def_func"`
	TimeWithTimeZoneDefFuncUniqueCheck  sql.NullTime `db:"time_with_time_zone_def_func_unique_check"`
}
