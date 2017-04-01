package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
	升级版,连接改为长连接，处理客户端发送的请求。调用客户端需要自己编写。
*/
func main() {
	service := ":1222"
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
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) //设置2分钟超时，当2分钟内客户端无请求发送时，关闭连接。
	request := make([]byte, 128)                          // request在创建时需要指定一个最大长度以防止flood attack
	defer conn.Close()                                    // 服务器端断开连接，客户端会接受到 An existing connection was forcibly closed by the remote host.的错误
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			fmt.Println("11111")
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			fmt.Println("22222")
			conn.Write([]byte(daytime))

		}

		request = make([]byte, 128) //每次读取到请求处理完毕后，需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
