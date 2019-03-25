package main

import (
	"fmt"
	"time"
)

//3种定时方法
func main() {
	/*fmt.Println("当前时间：",time.Now())
	myTimer:=time.NewTimer(time.Second*2)
	nowTime:=<-myTimer.C
	fmt.Println("现在时间：",nowTime)*/

	/*time.Sleep(time.Second)*/

	//time.After
	/*ti:=time.After(time.Second*2)
	fmt.Println(<-ti)*/

	//定时器重置和停止
	myTime := time.NewTimer(time.Second * 10)
	myTime.Reset(time.Second)
	go func() {
		<-myTime.C
		fmt.Println("子go程，定时完毕")
	}()

	for {

	}
}
