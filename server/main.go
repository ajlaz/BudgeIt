package main

import (
	"fmt"

	"github.com/ajlaz/BudgeIt/api"
	"github.com/ajlaz/BudgeIt/postgres"
)

func main() {
	fmt.Println("Hello, World!")
	db, err := postgres.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	postgres.Init(db)
	defer db.Close()
	a := api.New(db)
	if err := a.Run(); err != nil {
		fmt.Println(err)
		return
	}

}
