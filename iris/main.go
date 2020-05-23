package main

import (
	"log"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	v := app.Party("/apis")
	v.Get("/roles/{id:path}", func(ctx iris.Context) {
		log.Println(ctx.Params().Get("id"))
	})

	app.Run(iris.Addr(":8888"))
}
