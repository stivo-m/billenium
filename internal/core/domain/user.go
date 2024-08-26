package domain

import "time"

type User struct {
	Base
	Name            string    `gorm:"size:255;not null;unique" json:"name"`
	Email           string    `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerifiedAt time.Time `gorm:"default:null" json:"email_verified_at"`
}
