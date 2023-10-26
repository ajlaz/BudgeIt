package models

type User struct {
	ID        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

type UserConfig struct {
	UserId      int64   `db:"user_id" json:"user_id"`
	MonthlyGoal float64 `db:"monthly_goal" json:"monthly_goal"`
}
