package main

import "fmt"

func main() {
	a := 1
	var p *int = &a //定义一个int类型的指针
	fmt.Println(*p) //*p取得是a的值，如果直接打印p则是一个地址。
}
