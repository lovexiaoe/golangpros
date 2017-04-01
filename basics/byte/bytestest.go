package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fieldstest()
	fmt.Println("---------------")
	indexRuneTest()
	fmt.Println("---------------")
	titleTest()
	fmt.Println("---------------")
	ToTitleTest()
	fmt.Println("---------------")
	BufferWrite()
	fmt.Println("---------------")
}

//fields方法根据空格分割成多个slice，如果是全部是空格，则返回空的列表。
func fieldstest() {
	s := []byte("this is a! ")
	fields := bytes.Fields(s)
	fmt.Println(string(fields[0]))
	fmt.Println(string(fields[1]))
	fmt.Println(string(fields[2]))
}

//rune和int32一样，用于区分character值和integer值
func indexRuneTest() {
	s := []byte("02this is3 a! 2")

	fmt.Println(s)
	fmt.Println("byte1", s[1])
	index := bytes.IndexRune(s, 50)
	fmt.Println(index)
}

//Title返回每个byte的unicode字母。
func titleTest() {
	s := []byte("02好this is3 a! 2")
	news := bytes.Title(s)
	fmt.Println(len(news))
	fmt.Println(string(news[0]))
	fmt.Println(string(news[2]))
}

//ToTitle返回每个byte的unicode字母。和Title类似，没有看出差别
func ToTitleTest() {
	s := []byte("02好this is3 a! 2")
	news := bytes.ToTitle(s)
	fmt.Println(len(news))
	fmt.Println(string(news[0]))
	fmt.Println(string(news[2]))
}

func BufferWrite() {
	var b bytes.Buffer //buffer 不需要初始化
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "word!")
	b.WriteTo(os.Stdout)
}
