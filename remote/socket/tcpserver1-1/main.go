package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/*
	升级版，由原来的单任务变为并发处理。
*/
func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() // 服务器端断开连接，客户端会接受到 An existing connection was forcibly closed by the remote host.的错误
	daytime := time.Now().String()
	fmt.Println(daytime)
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
