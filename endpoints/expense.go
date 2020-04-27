package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	. "github.com/Prabandham/cost-tracker/objects"
)

type ExpenseParams struct {
	Amount        string `json:"expense_amount" binding:"required"`
	Description   string `json:"description"`
	ExpenseTypeId string `json:"expense_type_id" binding:""required"`
	AccountId     string `json:"account_id" binding:"required"`
}

func (e Endpoints) ListExpenses(c *gin.Context) {
	expenses := []Expense{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("user_id = ?", uid).Preload("Account").Preload("ExpenseType").Find(&expenses)
	c.JSON(http.StatusOK, expenses)
}

func (e Endpoints) CreateExpense(c *gin.Context) {
	expenseParams := ExpenseParams{}
	expense := Expense{}
	uid, _ := uuid.FromString(c.Param("user_id"))
	if err := c.ShouldBindJSON(&expenseParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}
	expense.UserID = uid
	expense.AccountID, _ = uuid.FromString(expenseParams.AccountId)
	expense.ExpenseTypeID, _ = uuid.FromString(expenseParams.ExpenseTypeId)
	expense.Amount, _ = strconv.ParseInt(expenseParams.Amount, 0, 64)
	expense.Description = expenseParams.Description
	e.Connection.FirstOrCreate(&expense, expense)
	e.Connection.Where("ID = ?", expense.ID).Preload("Account").Preload("ExpenseType").First(&expense)
	c.JSON(http.StatusCreated, expense)
}
