package postgres

import (
	"strings"

	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Insert(db *sqlx.DB, t *models.Transaction) error {
	if !validateDate(t.Date) {
		return models.ErrInvalidDate
	}
	_, err := db.Exec(`INSERT INTO transactions (user_id, amount, type, date) VALUES ($1, $2, $3, $4)`, t.UserID, t.Amount, t.Type, t.Date)
	return err
}

func validateDate(date string) bool {
	//12/03/2004
	if len(date) != 10 {
		return false
	}
	dateArr := strings.Split(date, "/")
	if len(dateArr) != 3 {
		return false
	}
	if len(dateArr[0]) != 2 || len(dateArr[1]) != 2 || len(dateArr[2]) != 4 {
		return false
	}
	return true
}
