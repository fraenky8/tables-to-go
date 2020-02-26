package dto

import (
	pg "github.com/lib/pq"
)

type TimeWithTimeZoneRef struct {
	TimeWithTimeZoneRef pg.NullTime `db:"time_with_time_zone_ref"`
}
