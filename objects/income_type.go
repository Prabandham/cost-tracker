package objects

import (
	uuid "github.com/satori/go.uuid"
)

// IncomeType can be defined by each user
// Ex Professional, Personal, HouseRent, Hobby, Part-Time-Work, Investment Returns, etc
type IncomeType struct {
	Base
	Name   string `sql:"index" json:"name" gorm:"unique;not null"`
	UserID uuid.UUID    `sql:"index" gorm:"type:uuid"`
	User   User
}
