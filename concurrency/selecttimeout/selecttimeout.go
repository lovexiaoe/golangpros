//select 结合timeout使用
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
		//如果不对c进行操作，则3秒后输出超时。
	case <-time.After(3 * time.Second):
		fmt.Println("TimeOut")
	}
}
