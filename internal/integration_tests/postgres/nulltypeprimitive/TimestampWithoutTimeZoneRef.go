package dto

type TimestampWithoutTimeZoneRef struct {
	TimestampWithoutTimeZoneRef *time.Time `db:"timestamp_without_time_zone_ref"`
}
