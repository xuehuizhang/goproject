package main

import (
	"fmt"
	"time"
)

func Producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
		fmt.Println("生产者生产：", i*i)
	}
	close(out)
}

func Consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到数据：", num)
		time.Sleep(time.Second)
	}
}

//生产者消费者模型     写入数据/消费数据/缓冲区
func main() {
	//写信人  生产者
	//邮递员  消费者
	//邮箱    缓冲区
	ch := make(chan int, 5)
	go Producer(ch)

	Consumer(ch)
}
