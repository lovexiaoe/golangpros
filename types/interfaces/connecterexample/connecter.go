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
	//Disconnect(a) //当实现接口的方法为T 和 *T
	Disconnect(&a) //当实现接口的方法为*T
	Disconnect2(a)

	//只有当接口为nil并不指向任何对象时，接口才为nil。
	var c interface{}
	fmt.Println(c == nil) //true;

	var p *int = nil
	c = p
	fmt.Println(c == nil) //false;

	interfaceAssign()

}

//可以将对象赋值给接口
/*
	将对象赋值给接口后，接口会成为包含两个字段的数据结构（这个结构叫interface value），
	第一个字段指向了一个内部表（iTable）和接口方法集，
	iTable中包含和存储值的类型信息（根据赋值给接口的类型不同，有value或者pointer，）
	第二个字段指向了存储值，这个存储值复制了被赋值的对象。所以对原对象的修改不会影响到interface value。
*/
func interfaceAssign() {

	var b Connecter
	//	b = PhoneConnecter{"valueTestConnecter"} //当实现接口的方法为T 和 *T
	b = &PhoneConnecter{"pointerTestConnecter"} //当实现接口的方法为*T
	b.Connect()
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

/**
	根据go官方文档，值类型的对象T,只能声明value receiver的方法作为方法集的一部分，
	而指针类型的对象*T既可以声明value receiver的方法也可以声明pointer receiver的方法作为方法集的一部分。
	Values 		Methods Receivers
	-----------------------------------------------
	T 			(t T)
	*T 			(t T) and (t *T)
	那将方法集和对象类型顺序反过来，就如下所示。
	Methods 	Receivers Values
	-----------------------------------------------
	(t T) 		T and *T
	(t *T) 		*T
**/

//当receiver为value类型时，value或者pointer类型的PhoneConnecter对象都会被认为成接口的实现者。
//func (pc PhoneConnecter) Connect() {
//	fmt.Println("connected.", pc.name)
//}

//当receiver是pointer类型时，只有pointer类型的PhoneConnecter对象才会被认为成接口的实现者。
func (pc *PhoneConnecter) Connect() {
	fmt.Println("connected.", pc.name)
}

func Disconnect(usb USB) {
	//这里usb是没有name属性的，用ok表达式判断是否为PhoneConnecter，然后打印出PhoneConnecter的name。
	if pc, ok := usb.(*PhoneConnecter); ok { //当实现接口的方法为*T
		//if pc, ok := usb.(PhoneConnecter); ok {//当实现接口的方法为T 和 *T
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
