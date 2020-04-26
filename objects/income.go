package objects

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Income struct {
	Base
	Amount         int       `sql:"index" json:"amount" gorm:"not null"`
	ReceivedOn     time.Time `json:"received_on"`
	UserID         uuid.UUID `sql:"index"`
	User           User
	IncomeSourceID uuid.UUID `sql:"index" json:"income_source_id`
	IncomeSource   IncomeSource
	AccountID      uuid.UUID `sql:"index" json:"account_id"`
	Account        Account
}
