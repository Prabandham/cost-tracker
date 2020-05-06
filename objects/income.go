package objects

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Income will hold all details pertaining to the income of a user.
// This will be linked to which account the income is going to and what the source of the income was
//
type Income struct {
	Base
	Account        Account
	AccountID      uuid.UUID `sql:"index" json:"account_id"`
	Amount         float64   `sql:"index" json:"amount" gorm:"not null"`
	Description    string    `gorm:"size: 255"`
	IncomeSource   IncomeSource
	IncomeSourceID uuid.UUID `sql:"index" json:"income_source_id`
	ReceivedOn     time.Time `json:"received_on"`
	User           User
	UserID         uuid.UUID `sql:"index"`
}

func (income *Income) BeforeCreate(scope *gorm.Scope) (err error) {
	account := Account{}
	db := scope.DB()
	db.Where("id = ?", income.AccountID).First(&account)
	newBalance := account.Balance + income.Amount
	db.Model(Account{}).Where("id = ?", account.ID).Update("balance", newBalance)
	return
}
