/*
	go中，IP的 定义如下。
	type IP []byte
*/

//运行时，加入一个ip作为参数：如window下：parseIP.exe 192.168.0.1

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-add\n", os.Args[0])
		os.Exit(0)
	}
	name := os.Args[1]
	addr := net.ParseIP(name) //把一个IPv4或者IPv6的地址转化成IP类型
	if addr == nil {
		fmt.Println("Invalid Address.")
	} else {
		fmt.Println("ip地址，The Address is", addr.String())
	}
	os.Exit(0)
}
