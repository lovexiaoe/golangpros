/*
	闭包的作用有1，模拟一个缓存变量，2，封装，外部不能直接访问到缓存变量，因为变量被定义在闭包作用域中。
	如下的x，在每次调用闭包函数后，x都会被改变，x这个变量不会给垃圾回收器回收。切外部不能直接访问到x。
	实现闭包一般是在非面向对象的语言中，在一个函数中定义一个函数，并返回这个函数。
*/

package main

import (
	"fmt"
)

func main() {
	f := closure(10)
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
}

func closure(x int) func(int) int {
	return func(y int) int {
		x = x + y
		return x
	}
}
