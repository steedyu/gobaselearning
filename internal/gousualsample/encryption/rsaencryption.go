package encryption

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"flag"
	"fmt"
	"errors"
	"encoding/base64"
)


func GenRsakeyDemo() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
	if err := GenRsaKey(bits); err != nil {
		fmt.Println("密钥文件生成失败！")
	}
	fmt.Println("密钥文件生成成功！")
}


//生成RSA密钥
func GenRsaKey(bits int) error {

	// 生成私钥文件
	//1 rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}


	//2 通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串（ASN.1是一种序列化方式，抽象反标记）
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	//3 将私钥字符串设置到pem格式快中
	// A Block represents a PEM encoded structure.
	//
	// The encoded form is:
	//    -----BEGIN Type-----
	//    Headers
	//    base64-encoded Bytes
	//    -----END Type-----
	// where Headers is a possibly empty sequence of Key: Value lines.
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	//4 通过pem将设置好的数据进行编码，并写入磁盘文件中
	privFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer privFile.Close()

	err = pem.Encode(privFile, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	//1 从得到的私钥对象中将公钥信息取出
	publicKey := &privateKey.PublicKey

	//2 通过x509标准将得到的rsa公钥序列化为字符串
	//虽然x509.MarshalPKIXPublicKey的参数是一个接口类型，其实对数据的类型是有要求的，和私钥是一样的
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	//3 将公钥字符串设置到pem格式快中
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	//4 通过pem将设置好的数据进行编码，并写入磁盘文件中
	pubFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	defer pubFile.Close()
	err = pem.Encode(pubFile, block)
	if err != nil {
		return err
	}
	return nil

}

/*
RSA加密解密
 */
func RsaEnDecryptDemo() {
	data, err := RsaPublicEncrypt([]byte("polaris@studygolang.com"))
	if err != nil {
		panic(err)
	}
	//这里要注意的一点是，由于加密后是字节流，直接输出查看会乱码，因此，为了便于语言直接加解密，这里将加密之后的数据进行base64编码。
	//fmt.Println("rsa encrypt :" + string(data))
	fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))

	origData, err := RsaPrivateDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 加密
func RsaPublicEncrypt(origData []byte) ([]byte, error) {
	// 2 得到的字符串解码
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	//3 使用x509将编码之后的公钥解析出来
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	//4 使用得到的公钥通过rsa进行数据加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaPrivateDecrypt(ciphertext []byte) ([]byte, error) {
	//2 得到的字符串解码
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//3 使用x509将编码之后的私钥解析出来
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//4 使用得到的私钥通过rsa进行数据解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}



var privateKey []byte = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDfEG3WBcHscVdlrY5cS4rNW4GXCD3roLTlGliBp8uI8zKmtJx3
Lk+VyZejH9zeaEUi+mKPMRpPtLZIWnVQYh0MeEj+rhqWG5yTN8d8g/KoxGU3whgF
RI/Sf2YCRmXRvCatm/aV3i83o70uQD4F5FxpZDdv5+ELl8yedNf+mtwztwIDAQAB
AoGBAJEHa4JFiAokvwAa0X5slzhhkGYUM74pZLO4Z2cVI55NENeWgkxyzcfDpFWo
97+a56iQRth2wnakNgfg2HmE8QDdb8IatiaffCcL7kROVbELIdYkYQzLy2FK7M/h
pW5GYLYg9Uue6mcEg0GVGUTNo7Ps5Kjegpzc+i5L02vLaPkBAkEA45I2VqifS/hr
ETbg6IucOVw5hbu2TBvwOaL9vJd1n3HxBuv2x5ks6x6vJO5LkhjfRi8d330lB3Pk
hLGNLN/GdwJBAPruFTFddPTyNZWAFaAf9E1NyfDnvA7u+AUB1RKQWI56rCiKdKuu
Ev6ibNcdokWeuB9yZhST2RvC8JMZrCJ2DMECQEYjn1nQOOCqXR1+I42o0eqf8R61
vzbv+XdaNAg3SkptTNNMUNAt9rk0yNiCFYqe3dn81aE3Kf2FC66WJqPpCHsCQQCe
Er9tAqe76p0Q2chFv/uBe0B8ry8L5UR+uwHEGQSAdQzg2R/YSueSWzXfab6gxvTM
cp+V1PGPCIXO1PxYFS/BAkEAlPYYFAtJIHrYm/ZBaUv2Pq4e2TE1rk4omgoFuAtt
EAJW+GmAgUsFzAga34bcMSMR8X+2svQ9X8tkeOkscqZTXQ==
-----END RSA PRIVATE KEY-----
`)

var publicKey []byte = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfEG3WBcHscVdlrY5cS4rNW4GX
CD3roLTlGliBp8uI8zKmtJx3Lk+VyZejH9zeaEUi+mKPMRpPtLZIWnVQYh0MeEj+
rhqWG5yTN8d8g/KoxGU3whgFRI/Sf2YCRmXRvCatm/aV3i83o70uQD4F5FxpZDdv
5+ELl8yedNf+mtwztwIDAQAB
-----END PUBLIC KEY-----
`)
