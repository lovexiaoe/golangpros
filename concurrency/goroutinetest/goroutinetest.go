package main

import (
	"fmt"
)

func main() {
	//goroutine的基本使用。
	//	go Go()
	//	//如果不进行睡眠，则main没有在Go执行完之前就退出了，导致GO执行了，但是程序GUI没有打印出来。
	//	time.Sleep(2 * time.Second)
	//----

	//channel进行进程间的通信，解决了上述需要睡眠的问题。
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!")
		//将true存入c
		c <- true
	}()
	//取出c,当main执行到这里时，会等到goroutine将值放入c中后再运行，所以会打印出Go,GO,GO
	<-c
	//---
}

func Go() {
	fmt.Println("GO GO GO!")
}
