package dto

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID         int            `db:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserID     int            `db:"user_id" gorm:"column:user_id;not null"`
	Email      string         `db:"email" gorm:"column:email;not null"`
	WebsiteURL sql.NullString `db:"website_url" gorm:"column:website_url"`
}
