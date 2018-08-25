package encryption

import (
	"crypto/cipher"
	"crypto/aes"
	"fmt"
	"encoding/base64"
)

/*
golang 的包中，AES的分组块长度只支持16个字节
C:\Go\src\crypto\aes\cipher.go
// The AES block size in bytes.
const BlockSize = 16

C:\Go\src\crypto\aes\cipher_amd64.go
func (c *aesCipherAsm) BlockSize() int { return BlockSize }

C:\Go\src\crypto\cipher\cipher.go
// A Block represents an implementation of block cipher
// using a given key. It provides the capability to encrypt
// or decrypt individual blocks. The mode implementations
// extend that capability to streams of blocks.
type Block interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Decrypt(dst, src []byte)
}
 */

func AesEnDeCryptDemo() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}


func AesEncrypt(origData, key []byte) ([]byte, error) {
	//1 创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	//2 对最后一个明文分组进行数据填充
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())

	//3 创建一个密码分组为链模式的，底层使用AES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	//4 加密连续的数据块
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	//1 创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	//2 创建一个密码分组为链模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	//3 解密连续的数据块
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)

	//4 去掉填充
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)

	return origData, nil
}



