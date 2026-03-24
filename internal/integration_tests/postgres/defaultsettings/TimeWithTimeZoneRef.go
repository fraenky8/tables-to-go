package dto

import (
	"database/sql"
)

type TimeWithTimeZoneRef struct {
	TimeWithTimeZoneRef sql.NullTime `db:"time_with_time_zone_ref"`
}
