package objects

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Expense is the actual expense incured by the person, which papps to the ExpenseType
type Expense struct {
	Base
	Amount        float64     `json:"amount" gorm:"not null"`
	SpentOn       time.Time   `json:"received_on"`
	Description   string      `gorm:"size: 255"`
	UserID        uuid.UUID   `sql:"index" gorm:"not null`
	User          User        `gorm:"PRELOAD`
	ExpenseTypeID uuid.UUID   `sql:"index gorm:"not null`
	ExpenseType   ExpenseType `gorm:"PRELOAD`
	AccountID     uuid.UUID   `sql:"index"`
	Account       Account     `gorm:"PRELOAD`
}

func (expense *Expense) BeforeCreate(scope *gorm.Scope) (err error) {
	account := Account{}
	db := scope.DB()
	db.Where("id = ?", expense.AccountID).First(&account)
	newBalance := account.Balance - expense.Amount
	db.Model(Account{}).Where("id = ?", account.ID).Update("balance", newBalance)
	return
}
