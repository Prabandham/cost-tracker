package endpoints

import (
    "net/http"

    "github.com/gin-gonic/gin"
    . "github.com/Prabandham/cost-tracker/objects"
    "github.com/satori/go.uuid"
)

func (e Endpoints) ListIncomeSources(c *gin.Context) {
    userParams := UserParams{}
    user := User{}
    incomeSources := []IncomeSource{}

    if err := c.ShouldBindUri(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
        return
    }

    uid, _ := uuid.FromString(userParams.UserId)
	e.Connection.Where("id = ?", uid).First(&user).Related(&incomeSources)
	c.JSON(http.StatusOK, incomeSources)
}