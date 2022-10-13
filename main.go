package main

import (
	"fmt"
	"todo_app/app/contollers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	contollers.StartMainServer()
}
