package objects

import uuid "github.com/satori/go.uuid"

type Account struct {
	Base
	Address          string  `json:"address" gorm:"size: 255"`
	Balance          float64 `json:"balance" binding:"required,min=1" gorm: not null`
	GlobalThreshold  float64 `json:"global_threshold"`
	IFSC             string  `json:"ifsc_code" binding:"required"`
	MonthlyThreshold float64 `json:"monthly_threshold"`
	Name             string  `json:"name" binding:"required"`
	User             User
	UserID           uuid.UUID `json:"user_id" gorm:"not null" sql:"index"`
}
