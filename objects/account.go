package objects

import uuid "github.com/satori/go.uuid"

type Account struct {
	Base
	Name    string    `json:"name" binding:"required"`
	Address string    `json:"address" gorm:"size: 255"`
	UserID  uuid.UUID `json:"user_id" binding:"required" gorm:"not null" sql:"index"`
	User    User
}
