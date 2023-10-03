package api

import (
	"errors"

	"github.com/ajlaz/BudgeIt/models"
	"github.com/ajlaz/BudgeIt/postgres"
	"github.com/gin-gonic/gin"
)

func (a *API) Login(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	temp, err := postgres.Login(&user, a.db)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, models.ErrIncorrectPassword) {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, temp.ID)

}
