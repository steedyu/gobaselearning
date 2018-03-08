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
此处事例 初始向量直接使用key充当了（实际项目中，最好别这么做）。

//密钥和向量必须为8位，否则加密解密都不成功。
 */
func DesEncrypt(origData, key []byte) ([]byte, error) {

	//这里使用的key
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())

	/*
	cipher.NewCBCEncrypter 会painc，需要进行相应处理
	 */

	//这里使用的是向量
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}


func PKCS5Padding(ciphertext []byte, blockSize int) []byte {

	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}



func DesDecrypt(crypted, key []byte) ([]byte, error) {

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)

	return origData, nil

}



func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


