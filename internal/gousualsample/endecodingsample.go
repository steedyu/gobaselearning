package gousualsample

import (
	"encoding/base64"
	"fmt"
)

func Base64Sample() {
	input := []byte("hello golang base64 快乐编程http://www.01happy.com +~")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(uDec))
}

func Base64Sample2() {
	Bytes, err := base64.StdEncoding.DecodeString("6bqm6L+q55S15rCU")
	fmt.Println(string(Bytes), err)
}