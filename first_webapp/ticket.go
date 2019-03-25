package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotterController struct {
	Ctx iris.Context
}

func newApp() *iris.Application  {
	app:=iris.New()
	mvc.New(app.Party("/",))
	return  app
}

func main()  {
	app:=newApp()
	app.Run(iris.Addr(":8080"))
}
