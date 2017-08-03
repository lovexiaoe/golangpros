//演示有缓存的channel，有缓存的channel是异步的，无缓存的channel是同步阻塞的。
//所以是否有缓存，读的顺序要先于写。

package main

import (
	"fmt"
)

//有缓冲，读先执行时，要等待缓存有写入，才执行读，所以可以打印出Go Go Go。
//func main() {
//	c := make(chan bool, 1)
//	go func() {
//		fmt.Println("GO GO GO!")
//		c <- true
//	}()
//	//在有缓冲区channel使用时，要读先于写，此程序会先执行下面语句发现c没有值后，会等goroutine执行。
//	<-c
//}

//有缓冲，写先执行，写不会等待读操作，所以写完时，main执行完成，所以打印不出Go Go Go。
//func main() {
//	c := make(chan bool, 1)
//	go func() {
//		fmt.Println("GO GO GO!")
//		<-c
//	}()
//	//在有缓冲区channel使用时，写先执行，写完后main执行完毕，不会等待goroutine的执行。
//	c <- true
//}

//channel无缓冲时，由于是同步阻塞，所以都能打印出Go Go Go
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		<-c
	}()
	c <- true
}
