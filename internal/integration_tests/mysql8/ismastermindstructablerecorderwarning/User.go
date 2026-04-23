package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type User struct {
	structable.Recorder

	ID         int            `db:"id"`
	UserID     int            `db:"user_id"`
	Email      string         `db:"email"`
	WebsiteURL sql.NullString `db:"website_url"`
}
