package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

func (e Endpoints) GetExpensesForMonth(c *gin.Context) {
	expenses := []Expense{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ? and month(spent_on) = 5", uid).Find(&expenses)
	c.JSON(http.StatusOK, expenses)
}

func (e Endpoints) GetExpensesGrouped(c *gin.Context) {
	uid := c.Param("user_id")
	expense := Expense{}
	result := expense.GetGroupedExpensesFor(e.Connection, uid)
	c.JSON(http.StatusOK, result)
}
