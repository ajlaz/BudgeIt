package postgres

import (
	"github.com/ajlaz/BudgeIt/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Init(db *sqlx.DB) {
	internal := db.DB
	gormdb, err := gorm.Open(pg.New(pg.Config{Conn: internal}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormdb.AutoMigrate(&models.User{}, &models.Transaction{})
}
