package dto

import (
	"database/sql"
)

type User struct {
	ID         int
	UserID     int
	WebsiteURL sql.NullString
}
