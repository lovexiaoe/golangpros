//方法的接收者，如下两个Print方法，他的接收者分别是type A和B，属于type的方法
package main

import "fmt"

func main() {
	a := new(A)
	a.Print()
	fmt.Println(a.Name)

	b := new(B)
	b.Print()
	fmt.Println(b.Name)

	var a1 A
	a1.Print()
	fmt.Println(a1.Name)

	var b1 B
	b1.Print()
	fmt.Println(b1.Name)
}

type A struct {
	Name string
}

type B struct {
	Name string
}

// value receiver 和 pointer receiver都可以被值类型和指针类型的对象调用，编译器会自动互相转换，
//但是通过接口对象调用方法时，pointer receiver方法只能被指针类型的接口对象调用，
//而value receiver方法既可以被指针类型的接口对象调用，也可以被值类型的接口对象调用。

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

//value receiver 不会对对象的状态造成改变，所以一般建议使用pointer receiver。
func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
