package objects

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const SigningKey string = "d74fb44d-1964-45ca-96de-a5edd74df6e8"

type User struct {
	Base
	Accounts             []Account
	Email                string `sql:"index" json:"email" gorm:"unique;not null" binding:"required,email"`
	EncryptedPassword    string `json:"-"`
	Expenses             []Expense
	IncomeSoruces        []IncomeSource
	Incomes              []Income
	Name                 string `sql:"index" json:"name" gorm:"not null" binding:"required"`
	Password             string `json:"password" gorm:"-" binding:"required,min=8,max=16"`
	PasswordConformation string `gorm:"-" json:"password_conformation" binding:"required,min=8,max=16,eqfield=Password"`
}

func (user *User) BeforeSave(scope *gorm.Scope) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	return scope.SetColumn("EncryptedPassword", password)
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	return err
}

func (user *User) GenerateJwtToken() (jwtToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["userid"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	token.Claims = claims
	jwtToken, err = token.SignedString([]byte(SigningKey))
	return jwtToken, err
}

func (user *User) IsValidSession(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errors.New("Token Expired")
		} else {
			return errors.New(err.Error())
		}
	} else {
		return errors.New(err.Error())
	}
}

func (user *User) GetIncomeAndExpenseSummary(db *gorm.DB, user_id string, year string) []map[string]string {
	results := make([]map[string]string, 0)
	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for _, month := range months {
		var income, expense string
		expenseRow := db.Table("expenses").Where("user_id = ?", user_id).Where("to_char(spent_on, 'Mon') = ?", month).Where("extract(year from spent_on) = ?", year).Select("sum(amount)").Row()
		expenseRow.Scan(&expense)
		incomeRow := db.Table("incomes").Where("user_id = ?", user_id).Where("to_char(received_on, 'Mon') = ?", month).Where("extract(year from received_on) = ?", year).Select("sum(amount)").Row()
		incomeRow.Scan(&income)
		result := map[string]string{"month": month, "year": year, "expense": expense, "income": income}
		results = append(results, result)
    }

	fmt.Println(results)
	return results
}
