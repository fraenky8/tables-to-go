package dto

import (
	"time"
)

type TimestampWithoutTimeZoneDefFuncUniqueCheckPk struct {
	TimestampWithoutTimeZoneDefFuncUniqueCheckPk time.Time `db:"timestamp_without_time_zone_def_func_unique_check_pk"`
}
