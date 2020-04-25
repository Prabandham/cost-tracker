package endpoints

import (
    "net/http"

    "github.com/gin-gonic/gin"
    . "github.com/Prabandham/cost-tracker/objects"
    "github.com/satori/go.uuid"
)

func (e Endpoints) ListIncomeSources(c *gin.Context) {
    incomeSources := []IncomeSource{}
    uid, _ := uuid.FromString(c.Param("user_id"))
	e.Connection.Where("id = ?", uid).Find(&incomeSources)
	c.JSON(http.StatusOK, incomeSources)
}