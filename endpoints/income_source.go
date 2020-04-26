package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

type IncomeSourceParams struct {
	Name string `json:name binding:"required"`
}

func (e Endpoints) ListIncomeSources(c *gin.Context) {
	incomeSources := []IncomeSource{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ?", uid).Find(&incomeSources)
	c.JSON(http.StatusOK, incomeSources)
}

func (e Endpoints) AddIncomeSource(c *gin.Context) {
	incomeSourceParams := IncomeSourceParams{}
	incomeSource := IncomeSource{}
	uuidParam := c.Param("user_id")
	if err := c.ShouldBindJSON(&incomeSourceParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
	uid, _ := uuid.FromString(uuidParam)
	incomeSource.Name = incomeSourceParams.Name
	incomeSource.UserID = uid
	e.Connection.FirstOrCreate(&incomeSource, incomeSource)
	c.JSON(http.StatusCreated, incomeSource)
}
