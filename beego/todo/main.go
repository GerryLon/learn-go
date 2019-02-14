package main

import (
	"github.com/GerryLon/learn-go/beego/todo/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo", &controllers.TodoController{}, "get:ListAll;delete:DelTodo")
	beego.Run()
}
