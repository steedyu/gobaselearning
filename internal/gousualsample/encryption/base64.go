package encryption

import (
	"encoding/base64"
	"errors"
	"fmt"
)

// 为了方便，声明该函数，省去错误处理
func mustDecode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// 该函数测试编解码
// enc为要测试的Encoding对象，str为要测试的字符串
func testEncoding(enc *base64.Encoding, str string) {
	// 编码
	encStr := enc.EncodeToString([]byte(str))
	fmt.Println(encStr)

	// 解码
	decStr := mustDecode(enc, encStr)
	if decStr != str {  // 编码后再解码应该与原始字符串相同
		// 这里判断如果不同，则panic
		panic(errors.New("unequal!"))
	}
}

func Base64Sample() {
	const testStr =  "guba.eastmoney.com"//"64D6EF4F-1D0C-404A-BDC7-8E6341421647" //"Go语言编程"

	// 测试StdEncoding，注意打印结果里的/为URL中的特殊字符，最后有一个padding
	testEncoding(base64.StdEncoding, testStr)    // 打印：R2/or63oqIDnvJbnqIs=

	// 测试URLEncoding，可以看到/被替换为_
	testEncoding(base64.URLEncoding, testStr)    // 打印：R2_or63oqIDnvJbnqIs=

	// 测试RawStdEncoding，可以看到去掉了padding
	testEncoding(base64.RawStdEncoding, testStr) // 打印：R2/or63oqIDnvJbnqIs

	// 测试RawURLEncoding，可以看到/被替换Wie_，并且却掉了padding
	testEncoding(base64.RawURLEncoding, testStr) // 打印：R2_or63oqIDnvJbnqIs
}

func GetOriCodeFromBase64Sample() {

	fmt.Println(mustDecode(base64.StdEncoding, "Y29tLmVhc3Rtb25leS5jb20uaXBob25l"))
	fmt.Println(mustDecode(base64.StdEncoding, "NjQ0NkVGNEUtMUQwQy00NTRBLUJEQzctOEU2MzQxNDI0NTQ3"))
	fmt.Println(mustDecode(base64.StdEncoding, "YzgzMTEwY2E1YmQwM2I5MjM5ZGJlNmRhMWMyZTE5ZjN8fDU4MTM4ODMzMDUwODg2OA=="))
	fmt.Println(mustDecode(base64.StdEncoding, "MjIyMjEtNDU0NjY3ODg="))

}

func OriCodeToBase64Sample() {
	testEncoding(base64.StdEncoding, "2233-45466788")
	testEncoding(base64.StdEncoding, "6446EF4E-1D0C-454A-BDC7-8E6341424547")   // NjQ0NkVGNEUtMUQwQy00NTRBLUJEQzctOEU2MzQxNDI0NTQ3
	testEncoding(base64.StdEncoding, "c83110ca5bd03b9239dbe6da1c2e19f3||581388330508868") //YzgzMTEwY2E1YmQwM2I5MjM5ZGJlNmRhMWMyZTE5ZjN8fDU4MTM4ODMzMDUwODg2OA==
	testEncoding(base64.StdEncoding, "com.eastmoney.com.iphone")    // Y29tLmVhc3Rtb25leS5jb20uaXBob25l
}
