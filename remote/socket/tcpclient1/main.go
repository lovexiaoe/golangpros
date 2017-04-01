//运行时，加入一个ip作为参数：如window下：tcpclient1.exe 192.168.0.1

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	/*
		func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
		ResolveTCPAddr方法的第一个参数为net，分别为"tcp4"、"tcp6"、"tcp"中的任意一个，
		分别表示TCP(IPv4-only),TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个).
		第二个参数addr表是域名或者ip地址。
	*/
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	/*
		func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
		net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
		laddr表示本机地址，一般设置为nil
		raddr表示远程的服务地址
	*/
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result))
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error1: %s", err.Error())
		os.Exit(1)
	}
}
