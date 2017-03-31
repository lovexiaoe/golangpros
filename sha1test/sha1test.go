package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	string1 := "jsapi_ticket=JSAPI_TICKET&noncestr=NONCESTR&timestamp=TIMESTAMP&url=URL"
	bytstr1 := []byte(string1)
	signature := sha1.Sum(bytstr1)

	fmt.Println(signature)
	s := fmt.Sprintf("%x", signature)
	fmt.Printf(s)

}
