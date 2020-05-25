package dto

import (
	"database/sql"
)

type DoublePrecisionRef struct {
	DoublePrecisionRef sql.NullFloat64 `db:"double_precision_ref"`
}
