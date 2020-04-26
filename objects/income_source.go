package objects

import (
	uuid "github.com/satori/go.uuid"
)

// IncomeSource can be defined by each user
// Ex Work, Rent, Hobby, Part-Time-Work, Investment Returns, etc
type IncomeSource struct {
	Base
	Name   string    `sql:"index" json:"name" gorm:"unique;not null"`
	UserID uuid.UUID `sql:"index" gorm:"type:uuid"`
	User   User
}
