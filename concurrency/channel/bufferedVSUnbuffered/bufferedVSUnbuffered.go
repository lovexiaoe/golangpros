package main

import (
	"fmt"
)

//有缓冲，可以打印出Go Go Go。
//func main() {
//	c := make(chan bool, 1)
//	go func() {
//		fmt.Println("GO GO GO!")
//		c <- true
//	}()
//	goroutine在buffered channel为空时读取会阻塞。所以这里会等待buffered channel有值写入后再读取。
//	等待goroutine完成，所以无论如何都会打印出Go Go GO!
//	<-c
//}

//不会打印Go Go Go。
//func main() {
//	c := make(chan bool, 1)
//	go func() {
//		fmt.Println("GO GO GO!")
//		<-c
//	}()
//	goroutine在buffered channel满了时写入会阻塞，之前没有写入，所以程序不会阻塞。main直接结束
//  来不及打印GO GO GO
//	c <- true
//}

//channel无缓冲时，读取和写入的动作必须同时完成，所以都能打印出Go Go Go
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		<-c
	}()
	c <- true
}
