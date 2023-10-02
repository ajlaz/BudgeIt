package postgres

import (
	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Insert(db *sqlx.DB, t *models.Transaction) error {
	_, err := db.Exec(`INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3)`, t.UserID, t.Amount, t.Type)
	return err
}
