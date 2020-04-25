package endpoints

import (
	"net/http"

	. "github.com/Prabandham/cost-tracker/objects"
	"github.com/gin-gonic/gin"
    "github.com/satori/go.uuid"
)

type UserParams struct {
	UserId string `uri:user_id binding:"required,uuid"`
}

type IncomeTypeParams struct {
    Name string `json:name binding:"required"`
	UserId string `json:user_id binding:"required,uuid"`
}

func (e Endpoints) ListIncomeTypes(c *gin.Context) {
	incomeTypes := []IncomeType{}
    uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ?", uid).Find(&incomeTypes)
	c.JSON(http.StatusOK, incomeTypes)
}

func (e Endpoints) AddIncomeType(c *gin.Context) {
	incomeTypeParams := IncomeTypeParams{}
	incomeType := IncomeType{}
	uuidParam := c.Param("user_id")
	if err := c.ShouldBindJSON(&incomeTypeParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
    uid, _ := uuid.FromString(uuidParam)
	incomeType.Name = incomeTypeParams.Name
	incomeType.UserID = uid
	e.Connection.FirstOrCreate(&incomeType, incomeType)
	c.JSON(http.StatusCreated, incomeType)
}
