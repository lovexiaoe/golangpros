package godoc_test

import (
	"fmt"

	"github.com/lovexiaoe/golangpros/basics/godoc"
)

// 以"包名_test"作为包名，
// 该go文件的名称必须以"_test"作为后缀，不然编译会报错。
//以Example作为方法名，在godoc中作为包的例子。

func Example() {
	oa := godoc.ObjectA{
		Name: "john",
	}
	fmt.Println("oa.Name", oa.GetOAName())

}

// Example+类型，可以作为类型的例子。

func ExampleObjectA() {
	fmt.Println("ObjectA 的例子命名为ExampleObjectA")
}

// Example+类型+_+方法名，可以作为方法的例子。

func ExampleObjectA_GetOAName() {
	fmt.Println("ObjectA.GetOAName 的例子命名为ExampleObjectA_GetOAName")
}
