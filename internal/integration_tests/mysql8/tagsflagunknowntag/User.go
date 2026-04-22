package dto

import (
	"database/sql"
)

type User struct {
	ID         int            `db:"id" json:"id"`
	UserID     int            `db:"user_id" json:"user_id"`
	Email      string         `db:"email" json:"email"`
	WebsiteURL sql.NullString `db:"website_url" json:"website_url"`
}
