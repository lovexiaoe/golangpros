//断定多个goroutine的执行完毕，方法一是使用缓存channel,方法二使用WaitGroup，waitGroup相当于定义了一个任务池。

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//方法一，缓存channel
//func main() {
//	//显式定义了逻辑处理器的个数，
//	//NumCPU取得当前机器的cpu核数。
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	c := make(chan bool, 10)
//	for i := 0; i < 10; i++ {
//		go Go(c, i)
//	}
//	//从缓冲channel中取10次。这样保证10个goroutine执行完后，main才退出。
//	for i := 0; i < 10; i++ {
//		<-c
//	}
//}

//func Go(c chan bool, index int) {
//	a := 1
//	for i := 0; i < 10000000; i++ {
//		a += i
//	}
//	fmt.Println(index, a)
//	//每个goroutine都会向缓冲channel中写入。
//	c <- true
//}

//方法2,WaitGroup方式
func main() {
	//显式定义了逻辑处理器的个数，
	//NumCPU取得当前机器的cpu核数。
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	//等待wg的任务池完成任务。
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//任务执行标记。
	wg.Done()
}
