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
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid Address.")
	} else {
		fmt.Println("ip地址，The Address is", addr.String())
	}
	os.Exit(0)
}
