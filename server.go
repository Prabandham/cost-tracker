package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Prabandham/cost-tracker/db"
	"github.com/Prabandham/cost-tracker/endpoints"
)

func main() {
	db := db.GetConnection()
	db.SetLogger()
	db.MigrateModels()

	server := gin.Default()

	loadRoutes(server, db)

	server.Run(":8080")
}

func loadRoutes(server *gin.Engine, db *db.Db) {
	// Expense Types
	endpoints := endpoints.Endpoints{Connection: db.Connection}

	// Public routes
	server.GET("/", endpoints.PingHandler)

	// Expense Type routes
	server.GET("/expense_types", endpoints.ListExpenseTypes)
	server.POST("/expense_type", endpoints.CreateExpenseType)

	// Users
	server.POST("/register", endpoints.RegisterUser)
	server.GET("/find_user", endpoints.FindUserByEmail)
	server.POST("/login", endpoints.Login)
	server.GET("/validate_session", endpoints.ValidateSession)
}
