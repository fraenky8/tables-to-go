package dto

import (
	"database/sql"
)

type TextTable struct {
	T            sql.NullString `db:"t"`
	TextNn       string         `db:"text_nn"`
	TextUnique   sql.NullString `db:"text_unique"`
	TextCheck    sql.NullString `db:"text_check"`
	TextRef      sql.NullString `db:"text_ref"`
	TextDefConst sql.NullString `db:"text_def_const"`
	TextPk       sql.NullString `db:"text_pk"`
}
