package dto

import (
	"database/sql"
)

type DecimalRef struct {
	DecimalRef sql.NullFloat64 `db:"decimal_ref"`
}
