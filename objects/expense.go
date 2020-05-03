package objects

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Expense is the actual expense incured by the person, which papps to the ExpenseType
type Expense struct {
	Base
	Amount        float64   `json:"amount" gorm:"not null"`
	SpentOn       time.Time `json:"spent_on"`
	Description   string    `gorm:"size: 255"`
	UserID        uuid.UUID `sql:"index" gorm:"not null`
	User          User
	ExpenseTypeID uuid.UUID `sql:"index gorm:"not null`
	ExpenseType   ExpenseType
	AccountID     uuid.UUID `sql:"index"`
	Account       Account
}

func (expense *Expense) BeforeCreate(scope *gorm.Scope) (err error) {
	account := Account{}
	db := scope.DB()
	db.Where("id = ?", expense.AccountID).First(&account)
	newBalance := account.Balance - expense.Amount
	db.Model(Account{}).Where("id = ?", account.ID).Update("balance", newBalance)
	return
}

func (expense *Expense) GetGroupedExpensesFor(db *gorm.DB, user_id string) []map[string]string {
	results := make([]map[string]string, 0)
	queryString := `select to_char(spent_on,'Mon') as Month,
						extract(year from spent_on) as Year,
	   					sum("amount") as "Spent Amount"
						from expenses as ex
						where ex.user_id = ?
						group by 1,2`

	rows, err := db.Raw(queryString, user_id).Rows()
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var month, year, amount string
		rows.Scan(&month, &year, &amount)
		result := map[string]string{"month": month, "year": year, "amount": amount}
		results = append(results, result)
	}

	return results
}
