package dto

import (
	"database/sql"
)

// This is the user table.
// Contains account information.
type User struct {
	ID int `db:"id"`
	// This is the identifier of the user.
	UserID int `db:"user_id"`
	// This is the email of the user.
	// Use it for notifications.
	Email string `db:"email"`
	// This is the website URL of the user.
	// Optional field.
	WebsiteURL sql.NullString `db:"website_url"`
}
