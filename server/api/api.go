package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type API struct {
	db *sqlx.DB
	*gin.Engine
}

func New(db *sqlx.DB) *API {
	return &API{
		db:     db,
		Engine: gin.Default(),
	}
}
