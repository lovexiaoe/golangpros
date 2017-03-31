//select写入channel的例子。
//一个空的select是完全阻塞的，它既没有发送，也没有接收。

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		//使用select向c中定入0或者1。这种写入是随机的
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
