package main

import (
	"fmt"
)

//go中的字符串都是采用UTF-8编码，字符串是用一对双引号（""）或者反引号（``）括起来的。
//不赋值时，默认为空字符串。
func main() {
	var emptyString string = "asdfdf"
	fmt.Printf(emptyString)

	var s string = "hello"
	//在go中字符串是不可变的，如下会报错

	//	s[0] = 'c'

	//如果想改需要转化成byte数组
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)

	//字符串虽不能修改，但可以做切片操作。
	s = "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)

}
