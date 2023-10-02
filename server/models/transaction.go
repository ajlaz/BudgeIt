package models

type Transaction struct {
	ID     int64   `db:"id" json:"id"`
	UserID int64   `db:"user_id" json:"user_id"`
	Amount float64 `db:"amount" json:"amount"`
	Type   string  `db:"type" json:"type"`
	Date   string  `db:"date" json:"date"`
}
