package dto

import (
	"database/sql"
)

type User struct {
	Id         int            `db:"id"`
	UserId     int            `db:"user_id"`
	Email      string         `db:"email"`
	WebsiteUrl sql.NullString `db:"website_url"`
}
