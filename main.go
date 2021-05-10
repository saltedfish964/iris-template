package main

import (
	"github.com/kataras/iris/v12"
	"iris-template/router"
)

func main() {
	app := iris.New()

	router.Init(app)

	app.Listen(":8088")
}
