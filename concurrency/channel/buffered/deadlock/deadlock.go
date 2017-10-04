//这里演示一个buffered channel 阻塞导致WaitGroup死锁的例子，当程序一直往channel发送时，当
//channel被写满，怎陷入阻塞，会导致WaitGroup陷入死锁。
//解决的方法可以是用一个for range不断读取。或者不让channel陷入阻塞。
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	channel := make(chan int, 2)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("send %d to channel", i)
			channel <- i
			wg.Done()
		}()
	}
	wg.Wait()
}
