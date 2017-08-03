/*
	跳转有goto,break,continue,配合标签使用

*/

package main

import "fmt"

func main() {
LABEL1:
	for {
		for i := 1; i < 10; i++ {
			if i > 3 {
				break LABEL1 //会跳出外层死循环。
				//goto LABEL1 不会跳出外层死循环。
				//所以在使用goto时，将标签放在循环后边，避免死循环，劲量避免使用goto。
			}
		}
	}
	fmt.Println("OK")
}
