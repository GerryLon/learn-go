package main

import (
	"github.com/GerryLon/learn-go/beego/todo/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo", &controllers.TodoController{}, "get:ListAll;post:AddTodo")
	beego.Router("/todo/:id:int", &controllers.TodoController{}, "delete:DelTodo;post:ModifyTodo")

	beego.Run()
}
