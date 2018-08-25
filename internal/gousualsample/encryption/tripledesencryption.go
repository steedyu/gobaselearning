package encryption

import (
	"crypto/des"
	"crypto/cipher"
)

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	//1 创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	//2 对最后一个明文分组进行数据填充
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())

	//3 创建一个密码分组为链模式的，底层使用3DES加密的BlockMode接口
	/*
	这里第二个参数是初始化向量 这里同样向量长度要和分组块长度相等，
	虽然密钥长度是24byte，但是每个块是8byte(因为是3次加解密)
	 */
	blockMode := cipher.NewCBCEncrypter(block, key[:8])

	//4 加密连续的数据块
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	//1 创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	//2 创建一个密码分组为链模式的，底层使用3DES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key[:8])

	//3 解密连续的数据块
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)

	//4 去掉填充
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)

	return origData, nil
}
