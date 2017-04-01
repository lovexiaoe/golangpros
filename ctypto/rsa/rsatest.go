package main

import (
	"crypto/rand"
	"crypto/rsa"
	// "crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	// "flag"s
	// "fmt"
	"log"
	// "time"
)

var decrypted string

// func init() {
//     flag.StringVar(&decrypted, "d", "", "加密过的数据")
//     flag.Parse()
// }

func main() {
	data, err := RsaEncrypt([]byte("qqqqqq"))
	if err != nil {
		panic(err)
	}
	// fmt.Println("1 ", string(data))
	log.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	// // log.Printf("private len", len(privateKey))
	// // log.Printf("timenow", time.Now().Format("2006-01-02 15:05:04"))

	// // log.Printf("public len", len(publicKey))
	// log.Println(len("2015-11-27 02:53:38"))
	b, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(data))

	log.Println(string(b))
	if err != nil {
		log.Printf("%q\n", err)
		panic(err)
	}

	origData, err := RsaDecrypt(b)
	if err != nil {
		log.Printf("%q\n", err)
		panic(err)
	}
	log.Println("2 ", string(origData))

	//对java加密的结果进行解密
	bb, err := base64.StdEncoding.DecodeString("1q4b5epRa4l5pLNVRZ9JkPIctYGLVFkHXxB8LvOCrQfRMvxk0l4VaeIUOPvWIPm6JTN/v7Vf/tGWBIoSyHzeR84Z4sYHNNY5sKdFJh81WLxuRrWz1mNzVDJbk+7GklZ6jA71FRhChxb1/gi8yPBXgz2iLfzMNPMz91uWMzYAojs=")
	if err != nil {
		log.Println("222222")
		log.Printf("%q\n", err)
		panic(err)
	}
	log.Println(string(bb))
	bbstr, err := RsaDecrypt(bb)

	if err != nil {
		log.Println("3333333333")
		log.Printf("%q\n", err)
		panic(err)
	}
	log.Println("3 ", string(bbstr))
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDld03V78DgoDo2Di97OhpY2ZThSAm9kfI/XAJz5BviNP63avRF
mdX+4VjT5dWPNK4bYeCZ+Iewn31GM1XZbDYDEyRPNV42iFQCDkYarUHOFhpqCpEs
G7qMkxg8HxnuDNh540FHHvFzXAXPsOzOxNZpkqLS9G2yZ5IhFEqjrVLlawIDAQAB
AoGAObOmfwWrGtEv0if/CJ2zwmP0bDIRQPpSUFxywXG7EUcCRl0+z8G/bjh8fcxt
x3UX0wrpz84PUPrKJb0C+Ymciu9/4lItXLb4h76B0GVm63AX3il2T0Eukial7PRT
+hxTZfZV4pcOz5vrVvJYarsdn8IXE3faNou7XLxAJgPJHQECQQDxWrETH6VQ+ZT1
V6yj1cYsdeksY6UpbIw98bGlnW/P/clSCi8vc6V9rBCEM7HsHtcmifqfFblTwpOy
AXpCd75PAkEA82PvYCL/AVZ4LCtXDwZdDeDtNldlRtfkKoYsnUMakSY2Ie6g7joY
wxbYEKXFjYbv4Q4SQr0nthSEEc5YlGBcJQJBAIvvl2eNG569dp5hfRlo4wP4QX+Z
LrO72fw4XFW32JJxhP5qJT2QAc3Bq7na9zf+EaSor4T5ZYCo+lVlAevz3YUCQQDv
+ypshTUgsYy+KGGny+N2qr/Z4+RVLMupbjCRQzfvxFh9rpd5LUl7GowiJgGa4WCm
bERvD6kXLDVohSfr7PMNAkEAi3dYtlXSgvTULVTA/e547rIwFifBwaj779IFFmAF
Dpx2TvL6/7kOn4WCItv93VgzoYdpaehWpXciwY26M9oWvA==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDld03V78DgoDo2Di97OhpY2ZTh
SAm9kfI/XAJz5BviNP63avRFmdX+4VjT5dWPNK4bYeCZ+Iewn31GM1XZbDYDEyRP
NV42iFQCDkYarUHOFhpqCpEsG7qMkxg8HxnuDNh540FHHvFzXAXPsOzOxNZpkqLS
9G2yZ5IhFEqjrVLlawIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	// return rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, ciphertext, nil)
}
