/*
	Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,
	UDP缺少了对客户端连接请求的Accept函数。
*/

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1333"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	fmt.Println(string(buf[0:n]))
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
