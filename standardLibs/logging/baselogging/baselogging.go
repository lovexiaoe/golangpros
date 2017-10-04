// 基本的日志输出
//log包中的logger是goroutine安全的，意味着你可以用多个goroutine同时调用同一个logger的方法集。
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger.
	log.Println("message")

	// Panicln is Println() followed by a call to panic().
	log.Panicln("panic message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	log.Fatalln("fatal message")

}
