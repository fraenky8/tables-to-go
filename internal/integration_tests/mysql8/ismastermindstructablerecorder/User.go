package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
)

type User struct {
	ID         int            `stbl:"id,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	UserID     int            `stbl:"user_id"`
	WebsiteURL sql.NullString `stbl:"website_url"`

	structable.Recorder
}
