/*
	for配合range可以用于读取slice和map的数据
*/

package main

import (
	"fmt"
)

func main() {
	amap := make(map[int]string)
	amap[1] = "first"
	amap[2] = "second"
	for k, v := range amap {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
}
