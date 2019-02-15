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

// 请求格式：
// POST /todo {"Title": "xxx"}
func (c *TodoController) AddTodo() {
	item := types.TodoItem{}
	result := util.Response{}
	var err error
	var success bool
	// err=unexpected end of JSON input
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &item)
	if err != nil {
		fmt.Printf("err=%s", err)
		result = util.FailResponse(nil)
		goto End
	}

	// 添加数据
	success, err = models.AddTodo(&item)
	if !success {
		result = util.FailResponse(nil)
		goto End
	}
	result.Data = types.TodoItem{
		Id: item.Id,
	}
End:
	c.Data["json"] = result
	c.ServeJSON()
}

// POST /todo/:id {Done: true} 将id对应的todo的done修改为true
// 目前只支持修改Done属性
func (c *TodoController) ModifyTodo() {
	var item types.TodoItem
	var err error
	var id int
	var success bool
	var result util.Response

	idStr := c.Ctx.Input.Param(":id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		result = util.FailResponse(nil)
		result.Ret = 1
		goto End
	}
	item = types.TodoItem{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &item)
	if err != nil {
		result = util.FailResponse(nil)
		result.Ret = 2
		goto End
	}
	item.Id = id
	success, err = models.ModifyTodo(&item)
	if !success {
		result = util.FailResponse(nil)
		result.Ret = 3
		goto End
	}
	result = util.SuccessResponse(nil)
End:
	c.Data["json"] = result
	c.ServeJSON()
}
