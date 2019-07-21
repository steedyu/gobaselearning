package encryption

import (
	"crypto/md5"
	"fmt"
)

func Md5Demo() {
	/*
	MD5加密的结果是和使用.Net中自带方法结果是一致的
	 */
	str := "abc123"
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	fmt.Println(md5str1)
}
