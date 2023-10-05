package api

import (
	"strconv"

	"github.com/ajlaz/BudgeIt/postgres"
	"github.com/gin-gonic/gin"
)

func (a *API) GetNum(c *gin.Context) {
	userId := c.Param("userId")
	n, err := strconv.Atoi(c.Param("num"))

	transactions, err := postgres.GetLastN(userId, n, a.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}
