package api

import (
	"github.com/ajlaz/BudgeIt/models"
	"github.com/ajlaz/BudgeIt/postgres"
	"github.com/gin-gonic/gin"
)

func (a *API) Register(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	err := postgres.Register(&user, a.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}
