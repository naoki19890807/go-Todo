package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	//init関数を呼ぶため
	fmt.Println(models.Db)

	controllers.StartMainSserver()
}
