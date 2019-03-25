package main

import (
	"fmt"
	"regexp"
)

//匹配小数
func main()  {
	str:="hello world 2.3 4.55 54.3 hfd 4fd .fd .4 5.fgfd"
	//rex:=`[0-9]*\.[0-9]*`
	rex:=`\d+\.\d+`
	req:=regexp.MustCompile(rex)
	alls:=req.FindAllStringSubmatch(str,-1)
	fmt.Println(alls)
}

//匹配字符
/*func main()  {
	str:="hello world e3l e4l eMl"

	reg:=regexp.MustCompile(`e[^0-9a-z]l`) //编译解析

    arr:=reg.FindAllStringSubmatch(str,-1) //提取需要信息

    fmt.Println(arr)
}*/
