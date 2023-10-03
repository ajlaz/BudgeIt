package postgres

import (
	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Register(user *models.User, db *sqlx.DB) error {
	inDb := models.User{}

	err := db.Get(&inDb, `SELECT * FROM users WHERE username = $1 LIMIT 1`, user.Username)
	if err == nil {
		return models.ErrUserAlreadyExists
	}

	_, err = db.Exec(`INSERT INTO users (username, password, first_name, last_name) VALUES ($1, $2, $3, $4)`, user.Username, HashPassword(user.Password), user.FirstName, user.LastName)
	return err
}
