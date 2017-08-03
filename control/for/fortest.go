//go中的循环只有for，没有while。

package main

import "fmt"

func main() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("Over")
}
