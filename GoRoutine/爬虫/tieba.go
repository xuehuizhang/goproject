package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Get(url string) (result string,err error) {
	res,err1:=http.Get(url)
	if err1!=nil{
		err=err1
		return
	}

	defer res.Body.Close()
	buf:=make([]byte,4069)
	for{
		n,err2:=res.Body.Read(buf)
		if err2!=nil&&err2!=io.EOF{
			err=err2
			return
		}
		if n==0{
			break
		}
		result+=string(buf)
	}
	return
}

func Splider(i int,page chan int)  {
	fmt.Printf("正在爬取第%d页...\n",i)
	url:="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
	res,err:=Get(url)
	if err!=nil{
		fmt.Println("work error:",err)
		return
	}

	fs,err:=os.Create("第"+strconv.Itoa(i)+"页"+string(time.Now().UnixNano())+".html")
	if err!=nil{
		fmt.Println("Os Create Err:",err)
		return
	}
	fs.WriteString(res)
	page<-i
}

func Work(start,end int)  {
	page :=make(chan int)

	for i:=start;i<=end;i++{
		go Splider(i,page)
	}
	for i:=start;i<=end;i++{
		fmt.Printf("第%d页读取完成\n",<-page)
	}
}

func main()  {
	var start,end int
	fmt.Print("请输入起始页（>1）:")
	fmt.Scan(&start)
	fmt.Print("请输入终止页（>start）")
	fmt.Scan(&end)
	Work(start,end)
}
