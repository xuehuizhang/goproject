package main

import "fmt"

//defer数据结构  栈的形式  先进后出
//defer--配合recover
func main()  {
	deferCall()
}
func deferCall()  {
	/*defer func() {
		fmt.Println("001")
	}()
	defer func() {
		fmt.Println("002")
	}()
	defer func() {
		fmt.Println("003")
	}()

	panic("出错误")*/

	deferRecover()
}

func deferRecover()  {
	defer func() {
		if ok:=recover();ok!=nil{
			fmt.Println("recover 处理错误")
		}
	}()
	panic("error")
}