package main

import "fmt"

//关闭channel 使用close
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("子go结束")
	}()

	/*for{
		if num,ok:=<-ch;ok==true{
			fmt.Println("读到数据：",num)
		}else{
			n:=<-ch
			fmt.Println("已经关闭",n)
			break
		}
	}*/

	for num := range ch {
		fmt.Println(num)
	}
}
