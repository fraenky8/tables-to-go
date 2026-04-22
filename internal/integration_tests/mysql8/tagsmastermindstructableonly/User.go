package dto

import (
	"database/sql"
)

type User struct {
	ID         int            `stbl:"id,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	UserID     int            `stbl:"user_id"`
	Email      string         `stbl:"email"`
	WebsiteURL sql.NullString `stbl:"website_url"`
}
