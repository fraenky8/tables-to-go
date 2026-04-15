package dto

import (
	"database/sql"
)

type StrictTypes struct {
	StrictID       sql.NullInt64   `db:"strict_id"`
	StrictInt      sql.NullInt64   `db:"strict_int"`
	StrictReal     sql.NullFloat64 `db:"strict_real"`
	StrictText     sql.NullString  `db:"strict_text"`
	StrictBlob     sql.NullString  `db:"strict_blob"`
	StrictAny      sql.NullString  `db:"strict_any"`
	StrictNn       string          `db:"strict_nn"`
	StrictUnique   sql.NullFloat64 `db:"strict_unique"`
	StrictCheck    sql.NullInt64   `db:"strict_check"`
	StrictRef      sql.NullInt64   `db:"strict_ref"`
	StrictDefConst sql.NullString  `db:"strict_def_const"`
}
