package dto

import (
	"database/sql"
)

type PrefixUser struct {
	ID         int            `db:"id"`
	UserID     int            `db:"user_id"`
	Email      string         `db:"email"`
	WebsiteURL sql.NullString `db:"website_url"`
}
