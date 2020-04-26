package objects

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Expense is the actual expense incured by the person, which papps to the ExpenseType
type Expense struct {
	Base
	Amount        int       `sql:"index" json:"amount" gorm:"not null"`
	SpentOn       time.Time `json:"received_on"`
	Description   string    `gorm:"size: 255"`
	UserID        uuid.UUID `sql:"index" gorm:"not null`
	User          User
	ExpenseTypeID uuid.UUID `sql:"index gorm:"not null`
	ExpenseType   ExpenseType
	AccountID     uuid.UUID `sql:"index"`
	Account       Account
}
