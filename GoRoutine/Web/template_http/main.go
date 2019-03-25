package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template

type Person struct {
	Name string
	Age int
}

type Output struct {
	out string
}

func (o *Output)Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	o.out+=string(p)
	n=len(p)
	return n,nil
}

func UserInfo(w http.ResponseWriter,r *http.Request){
	p:=make(map[string]interface{})
	p["Title"]="个人网站"
	p["Name"]="张三"
	p["Age"]=19
	//p:=Person{"张三",16}
	//myTemplate.Execute(w,p)
	//ou:=&Output{}
	myTemplate.Execute(w,p)     //注意这里ou 是一个实现了Write方法的接口，完全可以自己重写，实现自己的业务
	/*io.WriteString(ou,"hello world")
	fmt.Println(ou.out)*/
}

func InitTemplate(name string)(err error){
	myTemplate,err=template.ParseFiles(name)
	if err!=nil{
		fmt.Println("template err",err)
		return
	}
	return
}


func main(){
	InitTemplate("E:/Project/src/GoRoutine/Web/template_http/index.html")
	http.HandleFunc("/user/info",UserInfo)
	http.ListenAndServe(":8080",nil)
}
