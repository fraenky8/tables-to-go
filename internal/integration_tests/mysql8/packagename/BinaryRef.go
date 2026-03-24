package models

import (
	"database/sql"
)

type BinaryRef struct {
	BinaryRef sql.NullString `db:"binary_ref"`
}
