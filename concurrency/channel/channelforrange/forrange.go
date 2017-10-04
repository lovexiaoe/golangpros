//channel 使用forrange和close进行控制
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		c <- true
		close(c)
	}()
	//for range进行迭代channel，直到使用close关闭,main进程才会结束。
	//注意在使用for range时，必须要配合close，并且close必须关闭成功，否则程序会死锁，奔溃退出。
	for v := range c {
		fmt.Println(v)
	}
}
