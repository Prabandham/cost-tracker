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
	Amount         int64     `sql:"index" json:"amount" gorm:"not null"`
	ReceivedOn     time.Time `json:"received_on"`
	UserID         uuid.UUID `sql:"index"`
	User           User
	IncomeSourceID uuid.UUID `sql:"index" json:"income_source_id`
	IncomeSource   IncomeSource
	AccountID      uuid.UUID `sql:"index" json:"account_id"`
	Account        Account
}

func (income *Income) BeforeCreate(scope *gorm.Scope) (err error) {
	account := Account{}
	db := scope.DB()
	db.Where("id = ?", income.AccountID).First(&account)
	newBalance := account.Balance + income.Amount
	db.Model(Account{}).Where("id = ?", account.ID).Update("balance", newBalance)
	return
}
