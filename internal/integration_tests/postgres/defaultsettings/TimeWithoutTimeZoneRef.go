package dto

import (
	"database/sql"
)

type TimeWithoutTimeZoneRef struct {
	TimeWithoutTimeZoneRef sql.NullTime `db:"time_without_time_zone_ref"`
}
