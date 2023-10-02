package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Init(db *sqlx.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		type TEXT NOT NULL,
		date TEXT NOT NULL
	)`)
}
