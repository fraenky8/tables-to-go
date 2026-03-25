package dto

type User struct {
	ID         int     `db:"id"`
	UserID     int     `db:"user_id"`
	WebsiteURL *string `db:"website_url"`
}
