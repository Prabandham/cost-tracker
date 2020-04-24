package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/Prabandham/cost-tracker/objects"
)

func (e Endpoints) ListExpenseTypes(c *gin.Context) {
	expenseTypes := []ExpenseType{}
	e.Connection.Order("name asc").Find(&expenseTypes)
	c.JSON(http.StatusOK, expenseTypes)
}

func (e Endpoints) CreateExpenseType(c *gin.Context) {
	expenseType := ExpenseType{}
	if err := c.ShouldBindJSON(&expenseType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e.Connection.FirstOrCreate(&expenseType, expenseType)
	c.JSON(http.StatusCreated, expenseType)

}
