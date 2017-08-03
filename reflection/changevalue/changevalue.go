//使用Elem函数，结合地址修改对象的值，
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

func main() {
	changeint()

	u := User{1, "Mike", 20}
	Set(&u)
	fmt.Println(u)
}

//修改类型为int的对象
func changeint() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)
}

//修改User对象的名称
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断是否为指针且可以被修改
		fmt.Println("can't change value")
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	//判断是否找到Name字段
	if !f.IsValid() {
		fmt.Println("Not found Name")
		return
	}
	//修改name字段。
	if f.Kind() == reflect.String {
		f.SetString("LILI")
	}

}
