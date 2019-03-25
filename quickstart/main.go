package main

import (
	"github.com/astaxie/beego/context"
	_ "quickstart/routers"
	_ "quickstart/models"
	"github.com/astaxie/beego"
	"runtime"
	"strconv"
)

func main() {
	beego.AddFuncMap("ShowPrePage",HandlePrePage)
	beego.AddFuncMap("ShowNextPage",HandleNextPage)
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
	beego.Run()
}

func HandlePrePage(i int)string{
	pageIndex:=i-1
	return strconv.Itoa(pageIndex)
}

func HandleNextPage(i int)string  {
	pageIndex:=i+1
	return strconv.Itoa(pageIndex)
}

func FilterUser(ctx *context.Context)  {
	userName := ctx.Input.Session("userName")
	if userName==nil && (ctx.Request.RequestURI != "/"&&ctx.Request.RequestURI!="/reg") {
		ctx.Redirect(302, "/")
	}
	runtime.GC()
}

