//type 可以定义自定义类型，如下我们将string类型定义为中文进行使用。
package main

import "fmt"

type 文本 string //声明中文为string类型的自定义类型，但是最好不要使用中文。

func main() {
	var b 文本 = "类型为中文类型。"
	fmt.Println(b)
	Int2string()
}

/**type newInt int,这里newInt和int在进行类型转换时，仍需显式转换，但是byte和rune确实是
uint8和int32的别名，可以直接转换，
*/
