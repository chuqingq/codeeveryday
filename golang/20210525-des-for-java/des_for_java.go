package main

import (
	"bytes"
	"crypto/des"
	"errors"
	"log"
    "fmt"
    "encoding/base64"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("程序开始....")
	key := []byte("tyxthjwt")
	data := []byte("1234567890")

	out, _ := MyEncrypt(data, key)
	fmt.Println("加密后:", out)
    b := base64.StdEncoding.EncodeToString(out)
	fmt.Println("base64后:", b)
	out, _ = MyDecrypt(out, key)
	fmt.Println("解密后:", string(out))
}
func MyEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}
func MyDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// https://blog.csdn.net/scybs/article/details/38279159
// 加密后: [75 246 116 33 215 68 101 124 104 88 130 5 112 249 160 249]
// base64后: S/Z0IddEZXxoWIIFcPmg+Q==
// 解密后: 1234567890