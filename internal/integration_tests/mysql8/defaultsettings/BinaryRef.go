package dto

import (
	"database/sql"
)

type BinaryRef struct {
	BinaryRef sql.NullString `db:"binary_ref"`
}
