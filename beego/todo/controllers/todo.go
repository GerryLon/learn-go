package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GerryLon/learn-go/beego/todo/models"
	"github.com/GerryLon/learn-go/beego/todo/types"
	"github.com/GerryLon/learn-go/beego/todo/util"
	"github.com/astaxie/beego"
	"strconv"
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

// DELETE /todo/id
// {data: nil, msg: "ok", ret: 0}
func (c *TodoController) DelTodo() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	// beego.Debug(idStr, id)

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

func (c *TodoController) AddTodo() {
	result := &types.TodoItem{}
	fmt.Printf("body=%s", c.Ctx.Input.RequestBody)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, result)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
