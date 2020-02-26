package dto

import (
	pg "github.com/lib/pq"
)

type TimeWithoutTimeZoneRef struct {
	TimeWithoutTimeZoneRef pg.NullTime `db:"time_without_time_zone_ref"`
}
