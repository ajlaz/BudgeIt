package api

import (
	"errors"

	"github.com/ajlaz/BudgeIt/models"
	"github.com/ajlaz/BudgeIt/postgres"
	"github.com/gin-gonic/gin"
)

func (a *API) Insert(c *gin.Context) {
	t := models.Transaction{}
	c.BindJSON(&t)
	err := postgres.Insert(a.db, &t)
	if err != nil {
		if errors.Is(err, models.ErrInvalidDate) {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, t)

}
