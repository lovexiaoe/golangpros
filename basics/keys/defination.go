//报名，一个程序只能有一个main包
package main

//导入其他包
import "fmt"

//常量的定义
const Pi = 3.14

//全局变量的定义
var name = "gopher"

//一般类型声明
type newType int

//结构的声明
type gopher struct{}

//接口的声明
type golang interface{}

func main() {
	Println("hello world")
}

/**
const 进行常量的定义
var   用于变量的定义，在函数内部可以省略，在函数外部一般用于全局变量的定义。
type  用于结构（struct）和接口（interface）的声明，
func  用于函数的声明。
*/
