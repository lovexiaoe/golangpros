package main

import "fmt"

func main() {
	a := 10
	if a := 1; a >= 1 {
		fmt.Println(a)
	}
	fmt.Println(a)
}

//打印结果 1,10
