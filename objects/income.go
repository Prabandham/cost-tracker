package objects

import (
	"time"
)

type Income struct {
	Base
	Amount     int       `sql:"index" json:"amount" gorm:"not null"`
	ReceivedOn time.Time `json:"received_on"`
	UserId     int       `sql:"index"`
	User       User
}
