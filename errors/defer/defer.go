//defer函数在函数体执行完之后，按照调用顺序的相反顺序逐个执行。

package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i= ", i)
		defer func() {
			fmt.Println("defer_closure i= ", i)
			//这里的i在匿名函数内部，是一个闭包。所以i是引用地址，所以在defer执行时i已经变为4了
		}()
		fs[i] = func() { fmt.Println("closure i= ", i) }
	}

	for _, f := range fs {
		f()
	}
}
