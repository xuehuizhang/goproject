package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGetDou(url string)(result string,err error){
	resp,err1:=http.Get(url)
	if err1!=nil{
		err=err1
		return
	}
	defer resp.Body.Close()

	buf :=make([]byte,4096)
	for{
		n,err2:=resp.Body.Read(buf)
		if n==0{
			return
		}
		if err2!=nil&&err2!=io.EOF{
			err=err2
			return
		}
		result+=string(buf[:n])
	}
	return
}

func SpadierDou(idx int,ch chan bool)  {
	url:="https://book.douban.com/top250?start="+strconv.Itoa((idx-1)*25)
	res,err:=HttpGetDou(url)
	if err!=nil{
		fmt.Println("Http Get error",err)
		return
	}
	//获取书名
	title:=regexp.MustCompile(`&#34; title="(.*?)"`)
	allTitles:=title.FindAllStringSubmatch(res,-1)
	fmt.Println(allTitles)
	fmt.Println(len(allTitles))
	info:=regexp.MustCompile(`<p class="pl">(.*?)</p>`)
	allInfos:=info.FindAllStringSubmatch(res,-1)
	fmt.Println(allInfos)
	issue:=regexp.MustCompile(`([0-9]*)人评价`)
	allIssue:=issue.FindAllStringSubmatch(res,-1)
	fmt.Println(allIssue)
	//往文件中写
	count:= len(allTitles)
	fs,err:=os.Create("第"+strconv.Itoa(idx)+"页.txt")
	if err!=nil{
		fmt.Println("Os Create error:",err)
		return
	}
	fmt.Println(len(allTitles),len(allInfos),len(allIssue))
	fs.WriteString("书名\t\t\t\t\t作者详情\t\t\t\t\t评论人数\n\r")
	for i:=0;i<count;i++{
		fs.WriteString(allTitles[i][1]+"\t\t\t\t"+allInfos[i][1]+"\t\t\t\t"+allIssue[i][1]+"\n\r")
	}
	fmt.Printf("第%d页爬取完成",idx)
	ch<-true
}

func toWorkDou(start,end int)  {

	ch:=make(chan bool)
	for i:=start;i<=end ;i++  {
		fmt.Printf("%d页开始爬取",i)
		go SpadierDou(i,ch)
	}
	for i:=start;i<=end ;i++  {
		<-ch
	}
}

func main()  {
	var start,end int
	fmt.Println("请输入爬取起始页（>0）")
	fmt.Scan(&start)
	fmt.Println("请输入爬取终止页(>start)")
	fmt.Scan(&end)
	toWorkDou(start,end)
}
