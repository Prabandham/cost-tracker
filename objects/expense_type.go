package objects

// ExpenseType is a global list of all common expenses that a person my have.
// This is not defined by the user, as each user can enter their own value and will be
// hard to do statistics on random data going forward. SO this is a master list that the user has to choose from
type ExpenseType struct {
	Base
	Name string `sql:"index" json:"name" gorm:"unique;not null" binding:"required"`
}
