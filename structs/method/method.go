//如下对int封装为另一个类型，并提供一个Increase方法
package main

import "fmt"

type TZ int

type A struct {
}

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
	//+=操作的左右两端类型必须匹配，虽然TZ底层类型是int，但是和TZ是不同类型，需要将int转换为TZ。
}
