package main

import "fmt"

// 通道 <-chan 只写
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("生产 %d \n", i)
		out <- i
	}
	// 关闭通道, 不关闭会报错
	close(out)
}

func consumer(in <-chan int) {
	for data := range in {
		fmt.Printf("消费 %d \n", data)
	}
}

func main() {
	channel := make(chan int)
	go producer(channel)
	consumer(channel)
}
