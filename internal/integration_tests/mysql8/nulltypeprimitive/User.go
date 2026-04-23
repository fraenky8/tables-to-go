package dto

type User struct {
	ID         int     `db:"id"`
	UserID     int     `db:"user_id"`
	Email      string  `db:"email"`
	WebsiteURL *string `db:"website_url"`
}
