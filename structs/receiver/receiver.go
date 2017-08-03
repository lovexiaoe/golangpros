//方法的接收者，如下两个Print方法，他的接收者分别是type A和B，属于type的方法
package main

import "fmt"

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}

type A struct {
	Name string
}

type B struct {
	Name string
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
