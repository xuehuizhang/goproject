package main

import (
	"fmt"
	"html/template"
	"os"
)
type Person struct {
	Name string
	Age int
}
func main()  {
	t,err:=template.ParseFiles("E:/Project/src/GoRoutine/Web/template/index.html") //加载模板
	if err!=nil{
		fmt.Println("parse file err",err)
		return
	}
	p:=Person{"张三",19}
	err=t.Execute(os.Stdout,p)   //渲染模板
	if err!=nil{
		fmt.Println("There was an err",err)
	}
}
