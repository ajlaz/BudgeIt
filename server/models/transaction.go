package models

type Transaction struct {
	ID        int64   `db:"id" json:"id"`
	AccountID int64   `db:"account_id" json:"account_id"`
	Amount    float64 `db:"amount" json:"amount"`
	Type      string  `db:"description" json:"description"`
}
