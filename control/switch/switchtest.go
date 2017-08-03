//switch语句可以使用任何类型或者表达式作为条件语句
//不需要写break,一旦条件符合自动终止
//如果需要继续执行下一个case，使用fallthrough语句,fallthrough不管下一个语句块条件是否成立，都会执行下一个case。

package main

import "fmt"

func main() {
	switch1()
	fmt.Println()
	switchFallthrough() //输出a>=0
}

func switch1() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2, 3:
		fmt.Println("a=2 or a=3")
	default:
		fmt.Println("None")
	}
}

func switchFallthrough() {
	a := 1 //可以将此句放到switch后面并加分号。则a成为switch内部的局部变量。
	switch {
	case a >= 0:
		fmt.Println("a>=0")
	case a >= 1:
		fmt.Println("a>=1")
	case a >= 2:
		fmt.Println("a>=2")
		fallthrough
		//如果a>=2,则会执行default
	default:
		fmt.Println("None")
	}
}
