package dto

import (
	"database/sql"
)

type User struct {
	ID         int
	UserID     int
	Email      string
	WebsiteURL sql.NullString
}
