package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s { //forrange 取得的值是无序的。
		//		go func() {
		//			//这里是一个闭包，凡不是传入参入，都是引用类型，所以v的打印结果不可预测。
		//			fmt.Println(v)
		//		}()
		go func(v string) {
			//v作为参数传递时，是外部变量的一个拷贝。
			fmt.Println(v)
		}(v)
	}
	select {} //防止主程序退出，
}
