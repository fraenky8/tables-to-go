package dto

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID         int            `db:"id"`
	UserID     int            `db:"user_id"`
	Email      string         `db:"email"`
	WebsiteURL sql.NullString `db:"website_url"`
}
