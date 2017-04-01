package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	getcryptorandom()
}

//该方法主要是对beego sesseion包中产生sessionid的理解。生成密码安全的随机数。
func getcryptorandom() {
	//rand.Reader是一个产生密码安全的伪随机生成器，根据操作系统的不同调用不同平台的随机方法。
	reader := rand.Reader
	b := make([]byte, 16)
	_, err := ReadAtLeast(reader, b, len(b))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(b))
}

//从r中读取至少min个字节到buf中，这个方法是取自io包中的ReadAtLeast方法。
func ReadAtLeast(r io.Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, fmt.Errorf("short buffer!")
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = fmt.Errorf("ErrUnexpectedEOF!")
	}
	return
}
