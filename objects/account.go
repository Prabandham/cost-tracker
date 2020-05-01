package objects

import uuid "github.com/satori/go.uuid"

type Account struct {
	Base
	Name             string    `json:"name" binding:"required"`
	Address          string    `json:"address" gorm:"size: 255"`
	Balance          float64   `json:"balance" binding:"required,min=1" gorm: not null`
	IFSC             string    `json:"ifsc_code" binding:"required"`
	UserID           uuid.UUID `json:"user_id" gorm:"not null" sql:"index"`
	User             User
	GlobalThreshold  float64 `json:"global_threshold"`
	MonthlyThreshold float64 `json:"monthly_threshold"`
}
