package objects

// IncomeType can be defined by each user
// Ex Professional, Personal, HouseRent, Hobbie, ParttimeWork, Investment Returns, etc
type IncomeType struct {
	Base
	Name   string `sql:"index" json:"name" gorm:"unique;not null"`
	UserID int    `sql:"index"`
	User   User
}
