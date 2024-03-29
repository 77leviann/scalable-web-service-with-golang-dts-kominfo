package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Message   string    `json:"message" gorm:"not null" valid:"required"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User
	Photo     *Photo
}

type CommentUpdate struct {
	Message string `json:"message" gorm:"not null" valid:"required"`
}


func (comment *Comment) Validate() error {
	if _, err := govalidator.ValidateStruct(comment); err != nil {
		return err
	}

	return nil
}

