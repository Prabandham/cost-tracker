package objects

import uuid "github.com/satori/go.uuid"

type Account struct {
	Base
	Name    string    `json:"name" binding:"required"`
	Address string    `json:"address" gorm:"size: 255"`
	Balance int64     `json:"balance" binding:"required" gorm: not null`
	IFSC    string    `json:"ifsc_code"`
	UserID  uuid.UUID `json:"user_id" binding:"required" gorm:"not null" sql:"index"`
	User    User
}

// TODO setup account threshold. This basically is the overall threshold below which the amount should not go below at all. Irrespective of Month or year.

// TODO setup account monthly threshold. This basically is the overall threshold for the current month.

// TODO setup validation that balance cannot go below negative for a given account
