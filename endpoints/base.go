package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Endpoints
type Endpoints struct {
	Connection *gorm.DB
}

//PingHandler
func (e Endpoints) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Pong !!"})
}
