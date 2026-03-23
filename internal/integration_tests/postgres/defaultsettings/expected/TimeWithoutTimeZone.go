package dto

import (
	"database/sql"
	"time"
)

type TimeWithoutTimeZone struct {
	TimeWithoutTimeZone                    sql.NullTime `db:"time_without_time_zone"`
	TimeWithoutTimeZoneNn                  time.Time    `db:"time_without_time_zone_nn"`
	TimeWithoutTimeZoneNnUnique            time.Time    `db:"time_without_time_zone_nn_unique"`
	TimeWithoutTimeZoneNnCheck             time.Time    `db:"time_without_time_zone_nn_check"`
	TimeWithoutTimeZoneNnRef               time.Time    `db:"time_without_time_zone_nn_ref"`
	TimeWithoutTimeZoneNnDefConst          time.Time    `db:"time_without_time_zone_nn_def_const"`
	TimeWithoutTimeZoneNnDefFunc           time.Time    `db:"time_without_time_zone_nn_def_func"`
	TimeWithoutTimeZoneNnUniqueCheck       time.Time    `db:"time_without_time_zone_nn_unique_check"`
	TimeWithoutTimeZoneUnique              sql.NullTime `db:"time_without_time_zone_unique"`
	TimeWithoutTimeZoneUniqueCheck         sql.NullTime `db:"time_without_time_zone_unique_check"`
	TimeWithoutTimeZoneUniqueRef           sql.NullTime `db:"time_without_time_zone_unique_ref"`
	TimeWithoutTimeZoneUniqueDefConst      sql.NullTime `db:"time_without_time_zone_unique_def_const"`
	TimeWithoutTimeZoneUniqueDefFunc       sql.NullTime `db:"time_without_time_zone_unique_def_func"`
	TimeWithoutTimeZoneCheck               sql.NullTime `db:"time_without_time_zone_check"`
	TimeWithoutTimeZoneCheckRef            sql.NullTime `db:"time_without_time_zone_check_ref"`
	TimeWithoutTimeZoneCheckDefConst       sql.NullTime `db:"time_without_time_zone_check_def_const"`
	TimeWithoutTimeZoneCheckDefFunc        sql.NullTime `db:"time_without_time_zone_check_def_func"`
	TimeWithoutTimeZoneRef                 sql.NullTime `db:"time_without_time_zone_ref"`
	TimeWithoutTimeZoneRefDefConst         sql.NullTime `db:"time_without_time_zone_ref_def_const"`
	TimeWithoutTimeZoneRefDefFunc          sql.NullTime `db:"time_without_time_zone_ref_def_func"`
	TimeWithoutTimeZoneRefUniqueCheck      sql.NullTime `db:"time_without_time_zone_ref_unique_check"`
	TimeWithoutTimeZoneDefConst            sql.NullTime `db:"time_without_time_zone_def_const"`
	TimeWithoutTimeZoneDefConstUniqueCheck sql.NullTime `db:"time_without_time_zone_def_const_unique_check"`
	TimeWithoutTimeZoneDefFunc             sql.NullTime `db:"time_without_time_zone_def_func"`
	TimeWithoutTimeZoneDefFuncUniqueCheck  sql.NullTime `db:"time_without_time_zone_def_func_unique_check"`
}
