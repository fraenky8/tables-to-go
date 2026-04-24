package dto

import (
	"database/sql"
)

type BlobTable struct {
	Col               sql.NullString `db:"col"`
	BlobDefConst      sql.NullString `db:"blob_def_const"`
	BlobDefFunc       sql.NullString `db:"blob_def_func"`
	BlobNn            string         `db:"blob_nn"`
	BlobNnCheckCmp    string         `db:"blob_nn_check_cmp"`
	BlobNnCheckFn     string         `db:"blob_nn_check_fn"`
	BlobNnDefConst    string         `db:"blob_nn_def_const"`
	BlobNnDefFunc     string         `db:"blob_nn_def_func"`
	BlobCheck         sql.NullString `db:"blob_check"`
	BlobCheckDefConst sql.NullString `db:"blob_check_def_const"`
	BlobCheckDefFunc  sql.NullString `db:"blob_check_def_func"`
}
