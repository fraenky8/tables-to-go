package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type User struct {
	structable.Recorder

	ID         int            `stbl:"id,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	UserID     int            `stbl:"user_id"`
	Email      string         `stbl:"email"`
	WebsiteURL sql.NullString `stbl:"website_url"`
}
