// db.go
package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Prabandham/cost-tracker/objects"
)

type Db struct {
	Connection *gorm.DB
}

var singleton *Db
var once sync.Once

// TODO get dialact as input and return correspoinding connection Ex. GetConnection("mysql"), GetConnection("psql")
func GetConnection() *Db {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_USER", "srinidhi"),
			getEnv("DB_NAME", "cost_tracker"),
			getEnv("DB_PASSWORD", ""),
		)
		fmt.Println(psqlInfo)
		db, err := gorm.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println(err)
			panic("Could not connect to database")
		}
		singleton = &Db{Connection: db}
	})
	return singleton
}

func (db *Db) SetLogger() {
	db.Connection.LogMode(true)
	db.Connection.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

// MigrateModels will automatically migrate models
// Pass list of model objects to this as sting and dynamically load the objects to migrate by doing a struct look up
func (db *Db) MigrateModels() {
	db.Connection.AutoMigrate(
		&objects.Account{},
		&objects.Expense{},
		&objects.ExpenseType{},
		&objects.Income{},
		&objects.IncomeSource{},
		&objects.User{},
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
