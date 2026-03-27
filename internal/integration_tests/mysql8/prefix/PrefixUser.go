package dto

import (
	"database/sql"
)

type PrefixUser struct {
	ID         int            `db:"id"`
	UserID     int            `db:"user_id"`
	WebsiteURL sql.NullString `db:"website_url"`
}
