package api

import (
	"github.com/ajlaz/BudgeIt/postgres"
	"github.com/gin-gonic/gin"
)

func (a *API) GetData(c *gin.Context) {
	userId := c.Param("userId")
	transactions, err := postgres.GetData(userId, a.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}
