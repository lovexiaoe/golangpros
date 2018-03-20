//defer函数在“函数体”执行完之后，按照调用顺序的相反顺序逐个执行。
//defer虽然在“函数体”执行完后，但是引用的变量值为defer出现时的值。
package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//先执行 calc("10", a, b)=3 再执行 calc("20", a, b)=2
//再执行 calc("2", a, calc("20", a, b)) --> calc("2", 0, 2)=2
//再执行 calc("1", a, calc("10", a, b)) --> calc("1", 1, 3)=4
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
