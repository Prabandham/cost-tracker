package objects

// IncomeSource is generally the source of Income ex from an organization
// Ex Office Name, Bank Name, Person who pays rent etc
type IncomeSource struct {
	Base
	Name         string `sql:"index" json:"name" gorm:"unique: not null"`
	IncomeTypeID int    `sql:"index"`
	IncomeType   IncomeType
	UserID       int `sql:"index"`
	User         User
}
