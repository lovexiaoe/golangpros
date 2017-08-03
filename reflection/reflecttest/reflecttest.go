//go反射主要用到TypeOf（取得接口的类型）方法和ValueOf（取得接口的值）方法
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("hello world.")
}

func main() {
	u := User{1, "JAMES", 12}
	info(u)

	AnonymousField()
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//判断传入的对象是否是结构对象。如果传入的是地址（如&u）等，则返回。
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("非 struct 对象")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	//打印字段名称和值。
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", f.Name, f.Type, val)
	}
	//打印方法名称和值
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

//取结构中的匿名字段
func AnonymousField() {
	m := Manager{User: User{1, "KOBE", 29}, title: "zhuren"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))                  //按照索引取得User字段,Anonymous属性为true。
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) //按照索引取得User中的Id字段。

}
