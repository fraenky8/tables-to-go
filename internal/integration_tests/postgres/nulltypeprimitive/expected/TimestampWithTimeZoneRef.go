package dto

type TimestampWithTimeZoneRef struct {
	TimestampWithTimeZoneRef *time.Time `db:"timestamp_with_time_zone_ref"`
}
