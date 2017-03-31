//select 用于多个channel的处理，和switch有些类似。执行多个channel信息的发送与接收
//下面这个例子并不是很完善，不能判断两个channel的同时关闭。

package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	//用于判断c1,c2的某一个的关闭状态，确保goroutine在main函数之前完成，这样才能看到打印结果。
	//这里无法用o做到判断c1和c2同时关闭。
	o := make(chan bool)
	go func() {
		//select 在执行完后会结束执行，一般通过下面的无限循环来一直接收channel信息
		for {
			fmt.Println("111")
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					fmt.Println("222")
					//channel关闭后，跳出for循环
					break

				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					fmt.Println("333")
					//channel关闭后，跳出for循环
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	//这里只需要关闭一个，o被goroutine写后main程序退出。
	close(c1)
	//close(c2)

	//o的读取，确保goroutine在main完成之前执行
	<-o
}
