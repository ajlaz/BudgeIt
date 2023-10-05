package postgres

import (
	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetData(id string, db *sqlx.DB) ([]models.Transaction, error) {
	rows, err := db.Query(`SELECT * FROM transactions WHERE user_id=$1`, id)
	if err != nil {
		return nil, err
	}
	transactions := []models.Transaction{}
	for rows.Next() {
		t := models.Transaction{}
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type, &t.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func GetLastN(id string, n int, db *sqlx.DB) ([]models.Transaction, error) {
	rows, err := db.Query(`SELECT * FROM transactions WHERE user_id=$1 ORDER BY date DESC LIMIT $2`, id, n)
	if err != nil {
		return nil, err
	}
	transactions := []models.Transaction{}
	for rows.Next() {
		t := models.Transaction{}
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type, &t.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
