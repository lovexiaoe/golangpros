//slice的坑1
package main

import (
	"fmt"
)

func sliceappend(s []int) {
	s = append(s, 3)
	//这里的append重新生成了slice，但是方法引用的旧地址还是没有变，所以s不变。
	//解决方法，显式的返回新生成的slice ：return s。
	fmt.Println(s)
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)
	sliceappend(s)
	fmt.Println(s)
}
