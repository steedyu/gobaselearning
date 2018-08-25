package encryption

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"io"
)

/*
Md5的两种方式来看，其他的sha1等哈希函数使用是相似的
 */

/*
添加数据只能一次
 */
func Md5Demo_1() {
	/*
	MD5加密的结果是和使用.Net中自带方法结果是一致的
	 */
	str := "abc123"
	data := []byte(str)

	//1 计算结果
	has := md5.Sum(data)
	//2 散列值格式化
	//go 转换为16进制的方式
	//1st
	//md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	//2nd
	md5str1 := hex.EncodeToString(has[:])

	fmt.Println(md5str1)
}

/*
使用接口对象的sum方法
可以多次添加数据
 */
func Md5Demo_2() {
	str := "abc123"

	//1 创建哈希接口
	hash := md5.New()
	//2 添加数据
	//1st
	io.WriteString(hash, str)
	//2nd
	//data := []byte(str)
	//hash.Write(data)

	//3 计算结果
	/*
	如果b传值，则是把前面添加数据hash运算，然后对b哈希运算，然后组合起来
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte
	 */
	has := hash.Sum(nil)

	//4 散列值格式化
	md5str1 := hex.EncodeToString(has[:])

	fmt.Println(md5str1)
}