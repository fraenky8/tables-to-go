package dto

import (
	"database/sql"
)

type User struct {
	ID          int            `db:"id"`
	User_ID     int            `db:"user_id"`
	Email       string         `db:"email"`
	Website_URL sql.NullString `db:"website_url"`
}
