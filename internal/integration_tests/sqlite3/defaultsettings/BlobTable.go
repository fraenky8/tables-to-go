package dto

import (
	"database/sql"
)

type BlobTable struct {
	B            sql.NullString `db:"b"`
	BlobNn       string         `db:"blob_nn"`
	BlobUnique   sql.NullString `db:"blob_unique"`
	BlobCheck    sql.NullString `db:"blob_check"`
	BlobRef      sql.NullString `db:"blob_ref"`
	BlobDefConst sql.NullString `db:"blob_def_const"`
	BlobPk       sql.NullString `db:"blob_pk"`
}
