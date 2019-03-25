package main

import "fmt"

func sendCh(ch chan<- int) {
	ch <- 4
}

func recvCh(ch <-chan int) {
	n := <-ch
	fmt.Println("读取到：", n)
}

func main() {
	ch := make(chan int) //双向

	//var sendCh chan <- int=ch
	//sendCh<-9  //单向channel不能读取 err

	/*var recvCh <- chan int=ch
	fmt.Println(<-recvCh) //单向channel 不能写入*/
	go func() {

		sendCh(ch)
	}()
	recvCh(ch)

}
