/**Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1
 */

package main

import (
	"fmt"
)

const (
	x = 'A'  // x == 65
	y = iota // y == 1 只要出现iota，按行递增，并不是出现iota时，才从0开始。所以这里y=1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

//下面是log包中定义日期格式的定义，第一行表示1<<0。
const (
	Ldate         = 1 << iota
	Ltime                         //隐式地赋值为1<<1
	Lmicroseconds                 //隐式地赋值为1<<2
	LstdFlags     = Ldate | Ltime //这是赋值会打断iota。
	Llongfile                     //打断后会默认赋上一行的值，Ldate | Ltime
	LShortfile                    //打断后会默认赋上一行的值，Ldate | Ltime
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	fmt.Println(Ldate, Ltime, Lmicroseconds, LstdFlags, Llongfile, LShortfile)
}
