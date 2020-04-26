package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

type AccountParams struct {
	Name string `json:"name" binding:"required"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
	uid, _ := uuid.FromString(uuidParam)
	account.Name = accountParams.Name
	account.UserID = uid
	e.Connection.FirstOrCreate(&account, account)
	c.JSON(http.StatusOK, account)
}
