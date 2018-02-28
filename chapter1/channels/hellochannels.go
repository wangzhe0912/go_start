package main

import (
	"fmt"
	"sync"
)

// 相当于声明了一个计数器
var wg sync.WaitGroup

func printer(ch chan int) {
	// 不断读取输入的数，实时处理
	for i := range ch {
		// 默认会在此处不断执行，因为通道没有关闭
		fmt.Println("Received %d ", i)
	}
	fmt.Println("done")
	wg.Done()
}

// main is the entry point for the program.
func main() {
	// 启动一个通道
	c := make(chan int)
	// 启动一个goroutine，相当于一个独立的进程
	go printer(c)
	// 计数器加1
	wg.Add(1)

	// 依次向通道中发送10个数
	for i := 1; i <= 10; i++ {
		c <- i
		fmt.Println("Send %d ", i)
	}
	// 关闭通道
	close(c)
	// 计数器等待，只有当WaitGroup为0时，才会退出。
	wg.Wait()
}
