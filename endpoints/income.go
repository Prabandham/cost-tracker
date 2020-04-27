package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

type IncomeParams struct {
	AccountId string `json:"account_id" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
	// ReceivedOn     time.Time `json:"received_on" binding:"required"`
	IncomeSourceId string `json:"income_source_id" binding:"required"`
}

func (e Endpoints) ListIncomes(c *gin.Context) {
	incomes := []Income{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ?", uid).Preload("Account").Preload("IncomeSource").Find(&incomes)
	c.JSON(http.StatusOK, incomes)
}

func (e Endpoints) CreateIncome(c *gin.Context) {
	incomeParams := IncomeParams{}
	income := Income{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	if err := c.ShouldBindJSON(&incomeParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
	income.UserID = uid
	income.AccountID, _ = uuid.FromString(incomeParams.AccountId)
	income.Amount, _ = strconv.ParseInt(incomeParams.Amount, 0, 64)
	income.IncomeSourceID, _ = uuid.FromString(incomeParams.IncomeSourceId)
	e.Connection.FirstOrCreate(&income, income)
	e.Connection.Where("ID = ?", income.ID).Preload("Account").Preload("IncomeSource").First(&income)
	c.JSON(http.StatusCreated, income)
}
