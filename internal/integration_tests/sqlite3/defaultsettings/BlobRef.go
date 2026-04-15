package dto

import (
	"database/sql"
)

type BlobRef struct {
	BlobRef sql.NullString `db:"blob_ref"`
}
