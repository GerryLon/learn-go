package controllers

import (
	"github.com/GerryLon/learn-go/beego/todo/models"
	"github.com/GerryLon/learn-go/beego/todo/util"
	"github.com/astaxie/beego"
)

type TodoController struct {
	beego.Controller
}

// api: GET /todo
// {data: [{}, {}], ret: 0, msg: "success"}
func (c *TodoController) ListAll() {
	result := util.SuccessResponse(models.GetAll())
	c.Data["json"] = result
	c.ServeJSON()
}

// DELETE /todo {id: 0}
// {data: nil, msg: "ok", ret: 0}
func (c *TodoController) DelTodo() {
	id, err := c.GetInt("id")
	result := util.Response{}

	if err != nil {
		result.Ret = 1
	} else {
		b, err := models.DelTodo(id)
		if err != nil {
			result.Ret = 2
		} else if b {
			result = util.SuccessResponse(nil)
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
