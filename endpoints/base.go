package endpoints

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

//Endpoints
type Endpoints struct {
	Connection *gorm.DB
}

func FormatErrors(err error) string {
	validationErrors := err.(validator.ValidationErrors)
	errors := make([]string, 3)
	for _, e := range validationErrors {
		validationKind := e.ActualTag()
		key := e.StructField()
		switch validationKind {
		case "required":
			errors = append(errors, fmt.Sprintf("%s %s", key, "is required"))
		case "email":
			errors = append(errors, fmt.Sprintf("%s %s", key, "is not of a valid format"))
		default:
			errors = append(errors, fmt.Sprintf("%s %s", key, "is invalid"))
		}
	}
	return strings.Join(errors, ",")
}

//PingHandler
func (e Endpoints) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Pong !!"})
}
