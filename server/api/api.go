package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type API struct {
	db *sqlx.DB
	*gin.Engine
	handlers *[]gin.HandlerFunc
}

func New(db *sqlx.DB) *API {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	a := &API{
		db:       db,
		Engine:   router,
		handlers: &[]gin.HandlerFunc{},
	}

	a.RegisterRoutes()
	return a
}
