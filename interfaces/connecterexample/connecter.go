package main

import "fmt"

type USB interface {
	Name() string
	Connecter //嵌套接口
}
type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
	Disconnect2(a)

	var b Connecter
	b = Connecter(a) //可以将对象赋值给接口。
	//将对象赋值给接口时，会发生对象的copy，而接口内部存储是指向这个复制品的指针，无法修改复制品的状态，也无法获取指针。
	//如下，对原对象的修改并不会影响到接口指向的复制品。
	a.name = "pc"
	b.Connect()

	//只有当接口为nil并不指向任何对象时，接口才为nil。
	var c interface{}
	fmt.Println(c == nil) //true;

	var p *int = nil
	c = p
	fmt.Println(c == nil) //false;
}

//接口调用不会做receiver的自动转换，指针类型的receiver 方法实现接口时，只有指针类型的对象实现了该接口。
//值类型的对象只有（t T) 结构的方法，虽然值类型的对象也可以调用(t *T) 方法，但这实际上是Golang编译器自动转化成了&t的形式来调用方法
//（t T）可以被T和*T调用，而（t *T）只能被*T调用。
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("connected.", pc.name)
}

func Disconnect(usb USB) {
	//这里usb是没有name属性的，用ok表达式判断是否为PhoneConnecter，然后打印出PhoneConnecter的name。
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect.", pc.name)
		return
	}
	fmt.Println("Unknown device Disconnect")
}

//这里可以传入任何的type(空接口)。使用type-switch来判断传入的类型。
func Disconnect2(usb interface{}) {
	switch v := usb.(type) { //这里的v相当于上面的pc，v取得usb的实际类型变量。
	case PhoneConnecter:
		fmt.Println("Disconnect.", v.name)
	default:
		fmt.Println("Unknown device Disconnect")
	}
}
