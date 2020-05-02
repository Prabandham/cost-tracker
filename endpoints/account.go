package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

type AccountParams struct {
	Name             string  `json:"name" binding:"required"`
	Address          string  `json:"address"`
	Balance          float64 `json:"balance" binding:"required,min=1"`
	IFSC             string  `json:"ifsc_code" binding:"required"`
	GlobalThreshold  float64 `json:"global_threshold"`
	MonthlyThreshold float64 `json:"monthly_threshold"`
}

func (e Endpoints) ListAccounts(c *gin.Context) {
	accounts := []Account{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ?", uid).Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}

func (e Endpoints) CreateAccount(c *gin.Context) {
	accountParams := AccountParams{}
	account := Account{}
	uuidParam := c.Param("user_id")
	if err := c.ShouldBindJSON(&accountParams); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
	uid, _ := uuid.FromString(uuidParam)
	account.Name = accountParams.Name
	account.Address = accountParams.Address
	account.Balance = accountParams.Balance
	account.IFSC = accountParams.IFSC
	account.GlobalThreshold = accountParams.GlobalThreshold
	account.MonthlyThreshold = accountParams.MonthlyThreshold
	account.UserID = uid
	e.Connection.FirstOrCreate(&account, account)
	c.JSON(http.StatusOK, account)
}
