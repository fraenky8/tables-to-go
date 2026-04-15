package dto

import (
	"database/sql"
)

type StrictTypes struct {
	StrictID       sql.NullString `db:"strict_id"`
	StrictInt      sql.NullString `db:"strict_int"`
	StrictReal     sql.NullString `db:"strict_real"`
	StrictText     sql.NullString `db:"strict_text"`
	StrictBlob     sql.NullString `db:"strict_blob"`
	StrictAny      sql.NullString `db:"strict_any"`
	StrictNn       string         `db:"strict_nn"`
	StrictUnique   sql.NullString `db:"strict_unique"`
	StrictCheck    sql.NullString `db:"strict_check"`
	StrictRef      sql.NullString `db:"strict_ref"`
	StrictDefConst sql.NullString `db:"strict_def_const"`
}
