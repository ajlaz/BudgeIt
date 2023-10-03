package postgres

import (
	"errors"

	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Login(user *models.User, db *sqlx.DB) (*models.User, error) {
	temp := models.User{}
	err := db.Get(&temp, `SELECT * FROM users WHERE username = $1 LIMIT 1`, user.Username)
	if err != nil {
		return nil, err
	}
	if !CheckPasswordHash(user.Password, temp.Password) {
		return nil, errors.New("Incorrect Password")
	}
	return &temp, nil
}
