package encryption

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
)

/*
DES

在用DES加密解密时，经常会涉及到一个概念：
块（block，也叫分组），模式（比如cbc），初始向量（iv），
填充方式（padding，包括none，用’\0’填充，pkcs5padding或pkcs7padding）。

多语言加密解密交互时，需要确定好这些。比如这么定：
采用3DES、CBC模式、pkcs5padding，初始向量用key充当；
另外，对于zero padding，还得约定好，对于数据长度刚好是block size的整数倍时，是否需要额外填充。
*/

func DesDemo() {
	key := []byte("sfe023f_")
	result, err := DesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

/*
origData 原始数据
key 密钥

说明：
此处事例 初始向量直接使用key充当了（实际项目中，最好别这么做）。
密钥和向量必须为8字节(byte) 即64个Bit，否则加密解密都不成功。
 */
func DesEncrypt(origData, key []byte) ([]byte, error) {

	//1 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2 对最后一个明文分组进行数据填充
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())



	//3 创建一个密码分组为链模式的，底层使用DES加密的BlockMode接口
	/*
	NewCBCEncrypter(b Block, iv []byte) BlockMode
	b  模式
	iv 初始向量　The length of iv must be the same as the Block's block size
	否则会panic
	 */
	blockMode := cipher.NewCBCEncrypter(block, key)
	//4 加密连续的数据块
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}


//crypted 加密数据
//key 密钥
func DesDecrypt(crypted, key []byte) ([]byte, error) {

	//1 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//2 创建一个密码分组为链模式的，底层使用DES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key)

	//3 解密连续的数据块
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)

	//4 去掉填充
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)

	return origData, nil
}

/*
填充最后一个分组函数
ciphertext 原始数据
blockSize 每个分组的数据长度

填充思路：
1 最后一个分组不够8字节，需要填充满8字节
2 如果最后一个分组是密钥长度的整数倍，在末尾添加一分组，分组中每个字节的值和当前分组的长度相等

 */
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {

	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool{
		return r == rune(0)
	})
}




