package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string)(result string,err error){
	res,err1:=http.Get(url)

	if err1!=nil{
		err=err1
		return
	}
	defer  res.Body.Close()

	buf:=make([]byte,4096)
	for{
		n,err2:=res.Body.Read(buf)
		if n==0{
			fmt.Println("读取网页完成")
			break
		}
		if err2!=nil&&err2!=io.EOF{
			err=err2
			return
		}

		result+=string(buf[:n])
	}
	return
}

func working(start,end int) {
	fmt.Printf("正在爬取第%d页到%d页....\n",start,end)

	//循环爬取每一页数据
	for i:=start;i<=end;i++{
		url:="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
		res,err:=HttpGet(url)
		if err!=nil{
			fmt.Println("HttpGet Eror:",err)
			continue
		}
		//保存到文件中
		f,err:=os.Create("第"+strconv.Itoa(i)+"页.html")
		if err!=nil{
			fmt.Println("Os Create Error:",err)
			continue
		}
		f.WriteString(res)
		f.Close()
	}
}

func main()  {
	//指定爬虫起始和终止页
	var start,end int
	fmt.Print("请输入爬取的起始页（>1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>start）:")
	fmt.Scan(&end)
	working(start,end)
}
