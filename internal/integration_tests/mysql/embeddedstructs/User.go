package dto

import (
	"database/sql"

	"github.com/Masterminds/structable"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	structable.Recorder

	ID         int            `db:"id" gorm:"column:id;primaryKey;autoIncrement;not null" stbl:"id,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`
	UserID     int            `db:"user_id" gorm:"column:user_id;not null" stbl:"user_id"`
	Email      string         `db:"email" gorm:"column:email;not null" stbl:"email"`
	WebsiteURL sql.NullString `db:"website_url" gorm:"column:website_url" stbl:"website_url"`
}
