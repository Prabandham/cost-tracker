package main

import (
    "time"
    "log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/Prabandham/cost-tracker/db"
	"github.com/Prabandham/cost-tracker/endpoints"
	. "github.com/Prabandham/cost-tracker/objects"
)

func main() {
	db := db.GetConnection()
	defer db.Connection.Close()
	db.SetLogger()
	db.MigrateModels()

	server := gin.Default()
	corsConfig := cors.New(cors.Config{
                  		AllowOrigins:     []string{"*"},
                  		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
                  		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
                  		ExposeHeaders:    []string{"Content-Length"},
                  		AllowCredentials: true,
                  		AllowOriginFunc: func(origin string) bool {
                  			return origin == "*"
                  		},
                  		MaxAge: 12 * time.Hour,
                  	})
	server.Use(corsConfig)
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
    		Key:         []byte("d74fb44d-1964-45ca-96de-a5edd74df6e8"),
    		Timeout:     time.Hour,
    		MaxRefresh:  time.Hour,
    		IdentityKey: "ID",
    		PayloadFunc: func(data interface{}) jwt.MapClaims {
    			if v, ok := data.(*User); ok {
    				return jwt.MapClaims{
    					"ID": v.ID,
    				}
    			}
    			return jwt.MapClaims{}
    		},
    		IdentityHandler: func(c *gin.Context) interface{} {
    			claims := jwt.ExtractClaims(c)
    			data := map[string]string{"ID": claims["userid"].(string)}
    			return data
    		},
    		Authorizator: func(data interface{}, c *gin.Context) bool {
    			return true
    		},
    		Unauthorized: func(c *gin.Context, code int, message string) {
    			c.JSON(code, gin.H{
    				"code":    code,
    				"message": message,
    			})
    		},
    		TokenLookup: "header: Authorization, query: token, cookie: jwt",
    		TokenHeadName: "Bearer",
    		TimeFunc: time.Now,
    	})

    	if err != nil {
    		log.Fatal("JWT Error:" + err.Error())
    	}


	loadRoutes(server, db, authMiddleware)

	server.Run(":8080")
}

func loadRoutes(server *gin.Engine, db *db.Db, authMiddleware *jwt.GinJWTMiddleware) {
	// Expense Types
	endpoints := endpoints.Endpoints{Connection: db.Connection}

	// Public routes
	server.GET("/", endpoints.PingHandler)
    server.POST("/register", endpoints.RegisterUser)
    server.POST("/login", endpoints.Login)

    auth := server.Group("/auth")
    auth.Use(authMiddleware.MiddlewareFunc())
    {
        // Expense Type routes
        auth.GET("/expense_types", endpoints.ListExpenseTypes)
        auth.POST("/expense_type", endpoints.CreateExpenseType)

        // Users
        auth.GET("/find_user", endpoints.FindUserByEmail)
        auth.GET("/validate_session", endpoints.ValidateSession)

        // IncomeTypes
        auth.GET("/income_types/:user_id", endpoints.ListIncomeTypes)
        auth.POST("/income_type/:user_id", endpoints.AddIncomeType)

        // IncomeSources
        auth.GET("/income_sources/:user_id", endpoints.ListIncomeSources)
    }
}
