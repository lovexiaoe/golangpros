//这里演示一个buffered channel 阻塞导致WaitGroup死锁的例子，当程序一直往channel发送时，当
//channel被写满，则陷入阻塞，会导致WaitGroup陷入死锁。
//解决的方法可以是用一个for range不断读取。不让channel陷入阻塞,如下注释中的代码。
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int, 2)
	defer close(channel)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		//此时定义一个局域变量，存储外层i的值。count会在goroutine启动时，从0到4递增。
		//goroutine直接引用i，由于i为外层变量，则的i值不可预期。
		var count = i
		go func() {
			fmt.Printf("send %d to channel\n", count)
			channel <- i
			wg.Done()
		}()
	}
	//	go func() {
	//		for range channel {
	//			<-channel
	//		}
	//	}()
	wg.Wait()
}
