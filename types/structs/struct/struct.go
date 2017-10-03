package main

import "fmt"

func main() {
	a := person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)

	b := &person{ //一般在初始化结构时，会直接加一个地址符号，实现引用传递。
		Name: "Mike",
		Age:  26,
	}

	fmt.Println(b.Age) //并且不用通过*b的方式取对象的属性。
	B(b)
	fmt.Println(*b)
}

type person struct {
	Name string
	Age  int
}

func A(per person) { //默认的struct传递的是值传递， 这比例并不会改变a的值
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) { //通过指针实现值传递， 在调用时传递对象的地址。
	per.Age = 15
	fmt.Println("B", per)
}
