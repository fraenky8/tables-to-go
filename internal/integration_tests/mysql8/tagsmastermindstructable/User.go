package dto

import (
	"database/sql"
)

type User struct {
	ID         int            `db:"id" stbl:"id,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	UserID     int            `db:"user_id" stbl:"user_id"`
	WebsiteURL sql.NullString `db:"website_url" stbl:"website_url"`
}
