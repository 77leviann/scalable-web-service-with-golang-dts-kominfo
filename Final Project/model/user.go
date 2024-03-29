package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"not null;unique" json:"email" valid:"required,email"`
	Username     string    `gorm:"not null;unique" json:"username" valid:"required"`
	Password     string    `gorm:"not null" json:"password" valid:"required,minstringlength(8)"`
	Age          int       `gorm:"not null" json:"age" valid:"required,range(17|100)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

type UserUpdate struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Username string `gorm:"not null;unique" json:"username" valid:"required"`
}

type LoginCredential struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(8)"`
}
