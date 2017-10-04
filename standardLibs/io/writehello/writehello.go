// Sample program to show how different functions from the
// standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main is the entry point for the application.
func main() {
	//	创建一个Buffer，Buffer实现了io.Writer接口使用io.Write方法向buffer中写入byte。
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	//用Fprintf方法向BUffer 追加 "World!"字符串。方法的第一个参数为io.Writer接口。
	//而buffer的指针实现了io.Writer接口，所以传入b的地址。
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to the stdout device.
	// Passing the address of a os.File value for io.Writer.
	// WriteTo方法的参数是io.Writer的接口。	os.File实现了io.Writer方法。
	b.WriteTo(os.Stdout)
}
