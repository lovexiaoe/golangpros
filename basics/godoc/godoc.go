// 要生成godoc，注释必须在被注解的对象上面，中间不能有空行。
// 每个package对应一个godoc页面。
//
// 这里是一个新的段落，用空的注释行分段。
//
//   代码格式前面加3个空格，
//   这里是代码格式。
//package上面的注释生成为Overview模块。
package godoc

import (
	"fmt"
)

// 这里是对const的注释。
const (
	C1 = "asdf" // 大写的exported常量，可以生成为godoc。
	c2 = 12     // 小写的unexported变量，不能生成为godoc。
)

// 这里是对struct ObjectA的注解
type ObjectA struct {
	Name string // 大写的exported变量和方法，可以生成为godoc。
	age  int    // 小写的exported变量和方法，不能生成为godoc。
}

type objectB struct {
	Name string
	age  int
}

// 方法GetOAName的注释
func (oa *ObjectA) GetOAName() string {
	return oa.Name
}

// Foo方法是exported方法，会生成为godoc。
func Foo() {
	fmt.Println("Foo!")
}

// ob的方法不能生成为godoc
func (ob *objectB) GetOAName() string {
	return ob.Name
}
